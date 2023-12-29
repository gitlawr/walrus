package dao

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinition"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinitionmatchingrule"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ResourceDefinitionMatchingRulesEdgeSave saves the edge matching rules of model.ResourceDefinition entity.
func ResourceDefinitionMatchingRulesEdgeSave(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.ResourceDefinition,
) error {
	if entity.Edges.MatchingRules == nil {
		return nil
	}

	// Default new items and create key set for items.
	var (
		newItems       = entity.Edges.MatchingRules
		newItemsKeySet = sets.New[string]()
	)

	for i := range newItems {
		if newItems[i] == nil {
			return errors.New("invalid input: nil relationship")
		}
		newItems[i].ResourceDefinitionID = entity.ID
		newItems[i].Order = i

		newItemsKeySet.Insert(newItems[i].Name)
	}

	// Add/Update new items.
	if len(newItems) != 0 {
		err := mc.ResourceDefinitionMatchingRules().CreateBulk().
			Set(newItems...).
			OnConflict(
				sql.ConflictColumns(
					resourcedefinitionmatchingrule.FieldResourceDefinitionID,
					resourcedefinitionmatchingrule.FieldName,
				),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil && !errors.Is(err, stdsql.ErrNoRows) {
			return err
		}
	}

	// Delete stale items.
	oldItems, err := mc.ResourceDefinitionMatchingRules().Query().
		Where(resourcedefinitionmatchingrule.ResourceDefinitionID(entity.ID)).
		All(ctx)
	if err != nil {
		return err
	}

	deletedIDs := make([]object.ID, 0, len(oldItems))

	for i := range oldItems {
		if newItemsKeySet.Has(oldItems[i].Name) {
			continue
		}

		deletedIDs = append(deletedIDs, oldItems[i].ID)
	}

	if len(deletedIDs) != 0 {
		_, err = mc.ResourceDefinitionMatchingRules().Delete().
			Where(resourcedefinitionmatchingrule.IDIn(deletedIDs...)).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetMatchingResourceDefinitionWith gets the matching resource definition given resource type and project name
// with the cache.
func GetMatchingResourceDefinitionWith(
	ctx context.Context,
	mc model.ClientSet,
	resourceType, projectName string,
	cache map[string]*model.ResourceDefinition,
) (*model.ResourceDefinition, error) {
	key := fmt.Sprintf("%s:%s", resourceType, projectName)

	if cache[key] != nil {
		return cache[key], nil
	}

	rd, err := GetMatchingResourceDefinition(ctx, mc, resourceType, projectName)
	if err != nil {
		return nil, err
	}

	if cache != nil {
		cache[key] = rd
	}

	return rd, nil
}

// GetMatchingResourceDefinition gets the matching resource definition given resource type and project name.
func GetMatchingResourceDefinition(
	ctx context.Context,
	mc model.ClientSet,
	resourceType, projectName string,
) (*model.ResourceDefinition, error) {
	definition, err := mc.ResourceDefinitions().Query().
		Modify(func(s *sql.Selector) {
			s.Where(
				sql.And(
					sql.EQ(resourcedefinition.FieldType, resourceType),
					sql.Or(
						sqljson.ValueContains(resourcedefinition.FieldApplicableProjectNames, projectName),
						sqljson.LenEQ(resourcedefinition.FieldApplicableProjectNames, 0),
					),
				),
			).
				SelectExpr(
					sql.Expr(fmt.Sprintf(`DISTINCT on(%s) *`, resourcedefinition.FieldType)),
				).
				OrderExpr(
					sql.Expr(
						fmt.Sprintf(
							`%s, jsonb_array_length(%s) DESC`,
							resourcedefinition.FieldType,
							resourcedefinition.FieldApplicableProjectNames,
						),
					),
				)
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return definition, nil
}
