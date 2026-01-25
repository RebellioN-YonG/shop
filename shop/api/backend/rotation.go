package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"Rotation" method:"post" summary:"ur 1st rotation api"`
	PicUrl string `json:"pic_url" v:"required#cannot be empty" dc:"pic_url"`
	Link   string `json:"link" v:"required#cannot be empty" dc:"link"`
	Sort   int    `json:"sort" dc:"sort"`
}

type RotationRes struct {
	// g.Meta `mime:"text/html" example:"string"`
	RotationId int `json:"rotation_id"`
}
