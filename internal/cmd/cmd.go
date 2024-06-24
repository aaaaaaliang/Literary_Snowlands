package cmd

import (
	"Literary_Snowlands/internal/consts"
	"Literary_Snowlands/internal/controller/backend"
	"Literary_Snowlands/internal/controller/frontend"
	"Literary_Snowlands/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  consts.PROJECT_NAME,
		Usage: consts.PROJECT_USAGE,
		Brief: consts.PROJECT_BRIEF,

		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(service.Middleware().Cors)
			s.Use(service.Middleware().MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					frontend.NewUserControllerV1(),
					frontend.NewUtilControllerV1(),
					frontend.NewBookControllerV1(),
					frontend.NewHomeControllerV1(),
					frontend.NewNewsControllerV1(),
					frontend.NewSearchControllerV1(),
				)
			})

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().JWTMiddleware,
				)
				group.Bind(
					backend.NewAuthorControllerV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
