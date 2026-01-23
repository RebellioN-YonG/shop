package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Response unique data structure
type Resp struct {
	Msg  string      `json:"msg" dc:"消息"`
	Data interface{} `json:"data" dc:"结果"`
}

// interfaces' data strcucture
type HelloReq struct {
	g.Meta `path:"/" method:"get"`
	Name   string `v:"required" dc:"姓名"`
	Age    int    `v:"required" dc:"年龄"`
}
type HelloRes struct {
	Content string `json:"content" dc:"返回结果"`
}

// router func
type Hello struct{}

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	res = &HelloRes{
		Content: fmt.Sprintf(
			"hello %s, age %d",
			req.Name,
			req.Age),
	}
	return
}

func ErrHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.Write("error occurs: ", err.Error())
		return
	}
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ErrHandler)
		group.Bind(
			new(Hello),
		)
	})
	s.SetPort(8000)
	s.Run()

}
