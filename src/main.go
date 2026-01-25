package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Hello struct{}

func (Hello) Say(r *ghttp.Request) {
	r.Response.Write("Hello World")
}

type CB struct{}

func (CB) Say(r *ghttp.Request) {
	r.Response.Write("wxcb!")
}

type DDL struct{}

func (DDL) Say(r *ghttp.Request) {
	r.Response.Write("ddl!")
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			new(Hello),
			new(CB),
			new(DDL),
		)
	})
	s.SetPort(8000)
	s.Run()
}
