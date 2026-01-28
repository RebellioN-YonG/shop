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
				// 公共中间件（不需要认证）
				group.Middleware(
					service.Middleware().ResponseHandler,
					service.Middleware().Ctx,
				)
				// 不需要认证的接口（登录/注册等）
				group.Bind(
					controller.Login,
					controller.Data,
				)
				// 需要认证的接口
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.Rotation,
						controller.Admin.Create,
						controller.Admin.Update,
						controller.Admin.Delete,
						controller.Admin.List,
					)
					group.ALLMap(g.Map{
						"admin/info": controller.Admin.Info,
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
