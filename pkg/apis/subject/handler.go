package subject

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/apis/subject/view"
	"github.com/seal-io/seal/pkg/casdoor"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/role"
	"github.com/seal-io/seal/pkg/dao/model/subject"
	"github.com/seal-io/seal/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/settings"
)

func Handle(mc model.ClientSet) Handler {
	return Handler{
		modelClient: mc,
	}
}

type Handler struct {
	modelClient model.ClientSet
}

func (h Handler) Kind() string {
	return "Subject"
}

// Basic APIs.

func (h Handler) Create(ctx *gin.Context, req view.CreateRequest) (view.CreateResponse, error) {
	entity := req.Model()

	err := h.modelClient.WithTx(ctx, func(tx *model.Tx) error {
		creates, err := dao.SubjectCreates(tx, entity)
		if err != nil {
			return err
		}

		entity, err = creates[0].Save(ctx)
		if err != nil {
			return err
		}

		if entity.Kind != types.SubjectKindUser ||
			entity.Domain != types.SubjectDomainBuiltin {
			return nil
		}

		// Create user from casdoor.
		var cred casdoor.ApplicationCredential

		err = settings.CasdoorCred.ValueJSONUnmarshal(ctx, tx, &cred)
		if err != nil {
			return err
		}

		err = casdoor.CreateUser(ctx, cred.ClientID, cred.ClientSecret,
			casdoor.BuiltinApp, casdoor.BuiltinOrg, req.Name, req.Password)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return model.ExposeSubject(entity), nil
}

func (h Handler) Delete(ctx *gin.Context, req view.DeleteRequest) error {
	entity, err := h.modelClient.Subjects().Query().
		Where(subject.ID(req.ID)).
		Select(
			subject.FieldID,
			subject.FieldKind,
			subject.FieldDomain,
			subject.FieldName,
			subject.FieldBuiltin).
		Only(ctx)
	if err != nil {
		return err
	}

	if entity.Builtin {
		return runtime.Error(http.StatusForbidden, "cannot delete builtin subject")
	}

	return h.modelClient.WithTx(ctx, func(tx *model.Tx) error {
		err = tx.Subjects().DeleteOne(entity).
			Exec(ctx)
		if err != nil {
			return err
		}

		if entity.Kind != types.SubjectKindUser ||
			entity.Domain != types.SubjectDomainBuiltin {
			return nil
		}

		// Delete user from casdoor.
		var cred casdoor.ApplicationCredential

		err = settings.CasdoorCred.ValueJSONUnmarshal(ctx, tx, &cred)
		if err != nil {
			return err
		}

		err = casdoor.DeleteUser(ctx, cred.ClientID, cred.ClientSecret,
			casdoor.BuiltinOrg, entity.Name)
		if err != nil {
			if !strings.HasSuffix(err.Error(), "not found") {
				return fmt.Errorf("failed to delete user from casdoor: %w", err)
			}
		}

		return nil
	})
}

func (h Handler) Update(ctx *gin.Context, req view.UpdateRequest) error {
	entity, err := h.modelClient.Subjects().Query().
		Where(subject.ID(req.ID)).
		Select(
			subject.FieldID,
			subject.FieldKind,
			subject.FieldDomain,
			subject.FieldName).
		Only(ctx)
	if err != nil {
		return err
	}

	return h.modelClient.WithTx(ctx, func(tx *model.Tx) error {
		updates, err := dao.SubjectUpdates(tx, req.Model())
		if err != nil {
			return err
		}

		err = updates[0].Exec(ctx)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			// Maybe nothing change but password.
		}

		if entity.Kind != types.SubjectKindUser ||
			entity.Domain != types.SubjectDomainBuiltin {
			return nil
		}

		// Update password.
		if req.Password == "" {
			return nil
		}

		var cred casdoor.ApplicationCredential

		err = settings.CasdoorCred.ValueJSONUnmarshal(ctx, tx, &cred)
		if err != nil {
			return err
		}

		err = casdoor.UpdateUserPassword(ctx, cred.ClientID, cred.ClientSecret,
			casdoor.BuiltinOrg, entity.Name, "", req.Password)
		if err != nil {
			return fmt.Errorf("failed to update user password to casdoor: %w", err)
		}

		return nil
	})
}

// Batch APIs.

var (
	queryFields = []string{
		subject.FieldDomain,
		subject.FieldName,
	}
	getFields = subject.WithoutFields(
		subject.FieldUpdateTime)
	sortFields = []string{
		subject.FieldID,
		subject.FieldCreateTime,
	}
)

func (h Handler) CollectionGet(
	ctx *gin.Context,
	req view.CollectionGetRequest,
) (view.CollectionGetResponse, int, error) {
	query := h.modelClient.Subjects().Query()
	if req.Kind != "" {
		query.Where(subject.Kind(req.Kind))
	}

	if queries, ok := req.Querying(queryFields); ok {
		query.Where(queries)
	}

	// Get count.
	cnt, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get entities.
	if limit, offset, ok := req.Paging(); ok {
		query.Limit(limit).Offset(offset)
	}

	if fields, ok := req.Extracting(getFields, getFields...); ok {
		query.Select(fields...)
	}

	if orders, ok := req.Sorting(sortFields, model.Desc(subject.FieldCreateTime)); ok {
		query.Order(orders...)
	}

	entities, err := query.
		// Allow returning without sorting keys.
		Unique(false).
		WithRoles(func(rq *model.SubjectRoleRelationshipQuery) {
			rq.Order(model.Desc(subjectrolerelationship.FieldCreateTime)).
				Where(subjectrolerelationship.ProjectIDIsNil()).
				Select(subjectrolerelationship.FieldRoleID).
				// Allow returning without sorting keys.
				Unique(false).
				WithRole(func(rq *model.RoleQuery) {
					rq.Select(role.FieldID)
				})
		}).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeSubjects(entities), cnt, nil
}

// Extensional APIs.
