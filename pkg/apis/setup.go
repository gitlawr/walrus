package apis

import (
	"context"
	"net/http"
	"time"

	"k8s.io/client-go/rest"

	"github.com/seal-io/seal/pkg/apis/catalog"
	"github.com/seal-io/seal/pkg/apis/cli"
	"github.com/seal-io/seal/pkg/apis/connector"
	"github.com/seal-io/seal/pkg/apis/cost"
	"github.com/seal-io/seal/pkg/apis/dashboard"
	"github.com/seal-io/seal/pkg/apis/debug"
	"github.com/seal-io/seal/pkg/apis/measure"
	"github.com/seal-io/seal/pkg/apis/perspective"
	"github.com/seal-io/seal/pkg/apis/project"
	"github.com/seal-io/seal/pkg/apis/role"
	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/apis/setting"
	"github.com/seal-io/seal/pkg/apis/subject"
	"github.com/seal-io/seal/pkg/apis/template"
	"github.com/seal-io/seal/pkg/apis/templatecompletion"
	"github.com/seal-io/seal/pkg/apis/ui"
	"github.com/seal-io/seal/pkg/apis/variable"
	"github.com/seal-io/seal/pkg/auths"
	"github.com/seal-io/seal/pkg/dao/model"
)

type SetupOptions struct {
	// Configure from launching.
	EnableAuthn bool
	ConnQPS     int
	ConnBurst   int
	// Derived from configuration.
	K8sConfig    *rest.Config
	ModelClient  *model.Client
	TlsCertified bool
}

func (s *Server) Setup(ctx context.Context, opts SetupOptions) (http.Handler, error) {
	// Prepare middlewares.
	account := auths.RequestAccount(opts.ModelClient, opts.EnableAuthn)
	throttler := runtime.RequestThrottling(opts.ConnQPS, opts.ConnBurst)
	rectifier := runtime.RequestShaping(opts.ConnQPS, opts.ConnQPS, 5*time.Second)
	wsCounter := runtime.If(
		// Validate websocket connection.
		runtime.IsBidiStreamRequest,
		// Maximum 10 connection per ip.
		runtime.PerIP(func() runtime.Handle {
			return runtime.RequestCounting(10, 5*time.Second)
		}),
	)
	i18n := runtime.I18n()

	// Initial router.
	apisOpts := []runtime.RouterOption{
		runtime.WithDefaultWriter(s.logger),
		runtime.WithDefaultHandler(ui.Index(ctx, opts.ModelClient)),
		runtime.SkipLoggingPaths(
			"/",
			"/cli",
			"/assets/*filepath",
			"/readyz",
			"/livez",
			"/metrics",
			"/debug/version"),
		runtime.ExposeOpenAPI(),
		runtime.WithResourceRouteAdviceProviders(provideModelClient(opts.ModelClient)),
		runtime.WithResourceAuthorizer(account),
	}

	apis := runtime.NewRouter(apisOpts...).
		Use(i18n)

	cliApis := apis.Group("")
	{
		r := cliApis
		r.Get("/cli", cli.Index())
	}

	measureApis := apis.Group("").
		Use(throttler)
	{
		r := measureApis
		r.Get("/readyz", measure.Readyz())
		r.Get("/livez", measure.Livez())
		r.Get("/metrics", measure.Metrics())
	}

	accountApis := apis.Group("/account").
		Use(rectifier, account.Filter)
	{
		r := accountApis
		r.Post("/login", account.Login)
		r.Post("/logout", account.Logout)
		r.Get("/info", account.GetInfo)
		r.Post("/info", account.UpdateInfo)
		r.Post("/tokens", account.CreateToken)
		r.Delete("/tokens/:token", account.DeleteToken)
		r.Get("/tokens", account.GetTokens)
	}

	resourceApis := apis.Group("/v1").
		Use(throttler, wsCounter, account.Filter)
	{
		r := resourceApis
		r.Resource(catalog.Handle(opts.ModelClient))
		r.Resource(connector.Handle(opts.ModelClient))
		r.Resource(cost.Handle(opts.ModelClient))
		r.Resource(dashboard.Handle(opts.ModelClient))
		r.Resource(perspective.Handle(opts.ModelClient))
		r.Resource(project.Handle(opts.ModelClient, opts.K8sConfig, opts.TlsCertified))
		r.Resource(role.Handle(opts.ModelClient))
		r.Resource(setting.Handle(opts.ModelClient))
		r.Resource(subject.Handle(opts.ModelClient))
		r.Resource(template.Handle(opts.ModelClient))
		r.Resource(templatecompletion.Handle(opts.ModelClient))
		r.Resource(variable.Handle(opts.ModelClient))
	}

	debugApis := apis.Group("/debug")
	{
		r := debugApis
		r.Get("/version", debug.Version())
		r.Get("/flags", debug.GetFlags())
		r.Group("").
			Use(runtime.OnlyLocalIP()).
			Get("/pprof/*any", debug.PProf()).
			Put("/flags", debug.SetFlags())
	}

	return apis, nil
}
