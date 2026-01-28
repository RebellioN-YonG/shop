package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"Rotation" method:"post" summary:"ur 1st rotation api"`
	PicUrl string `json:"pic_url" v:"required#cannot be empty" dc:"pic_url"`
	Link   string `json:"link" v:"required#cannot be empty" dc:"link"`
	Sort   int    `json:"sort" dc:"sort"`
}

type RotationRes struct {
	RotationId int `json:"rotation_id"`
}

type RotationDeleteReq struct {
	g.Meta     `path:"/rotation/delete" tags:"Rotation" method:"delete" summary:"delete a rotation"`
	RotationId uint `json:"rotation_id" v:"min:1#select a rotation to delete" dc:"rotation_id"`
}

type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta     `path:"/rotation/update" tags:"Rotation" method:"post" summary:"update a rotation"`
	RotationId uint   `json:"rotation_id" v:"min:1#select a rotation to update" dc:"rotation_id"`
	PicUrl     string `json:"pic_url" v:"required#cannot be empty" dc:"pic_url"`
	Link       string `json:"link" v:"required#cannot be empty" dc:"link"`
	Sort       int    `json:"sort" dc:"sort"`
}

type RotationUpdateRes struct {
	Id uint `json:"id"`
}

type RotationGetListCommonReq struct {
	g.Meta `path:"/rotation/list" tags:"Rotation" method:"get" summary:"get rotation list"`
	Sort   int `json:"sort" in:"query" dc:"sort"`
	CommonPaginationReq
}

type RotationGetListCommonRes struct {
	List  interface{} `json:"list" description:"rotation list"`
	Page  int         `json:"page" description:"page number"`
	Size  int         `json:"size" description:"size of page"`
	Total int         `json:"total" description:"total amount"`
}
