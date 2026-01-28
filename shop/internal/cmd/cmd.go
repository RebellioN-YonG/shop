package cmd

import (
	"context"
	"shop/internal/controller"
	"shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					// ghttp.MiddlewareHandlerResponse,
					service.Middleware().ResponseHandler,
					service.Middleware().Ctx,
					service.Middleware().Auth,
				)
				group.Bind(
					controller.Rotation, // Register Rotation
					controller.Admin,    // Register Admin
					controller.Login,
				)
				// group.Bind()
			})
			s.Run()
			return nil
		},
	}
)
