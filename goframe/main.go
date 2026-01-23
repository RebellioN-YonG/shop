package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	svr := g.Server()
	svr.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writef(
			"Hello %s, age %d",
			r.Get("name", "unknown").String(),
			r.Get("age", 0).Int(),
		)
	})
	svr.SetPort(8080)
	svr.Run()
}
