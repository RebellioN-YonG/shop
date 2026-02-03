package backend

import "github.com/gogf/gf/v2/frame/g"

type PositionReq struct {
	g.Meta    `path:"/position/add" tags:"Position" method:"post" summary:"ur 1st position api"`
	PicUrl    string `json:"pic_url" v:"required#cannot be empty" dc:"pic_url"`
	Link      string `json:"link" v:"required#cannot be empty" dc:"link"`
	Sort      int    `json:"sort" dc:"sort"`
	GoodsName string `json:"goods_name" v:"required#goods name cannot be empty" dc:"goods name"`
	GoodsId   uint   `json:"goods_id" v:"required#goods id cannot be empty" dc:"goods id"`
}

type PositionRes struct {
	PositionId int `json:"position_id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/position/delete" tags:"Position" method:"delete" summary:"delete a position"`
	Id     uint `json:"id" v:"min:1#select a position to delete" dc:"position_id"`
}

type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta     `path:"/position/update" tags:"Position" method:"post" summary:"update a position"`
	PositionId uint   `json:"position_id" v:"min:1#select a position to update" dc:"position_id"`
	PicUrl     string `json:"pic_url" v:"required#cannot be empty" dc:"pic_url"`
	Link       string `json:"link" v:"required#cannot be empty" dc:"link"`
	Sort       int    `json:"sort" dc:"sort"`
	GoodsName  string `json:"goods_name" v:"required#goods name cannot be empty" dc:"goods name"`
	GoodsId    uint   `json:"goods_id" v:"required#goods id cannot be empty" dc:"goods id"`
}

type PositionUpdateRes struct {
	Id uint `json:"id"`
}

type PositionGetListCommonReq struct {
	g.Meta `path:"/position/list" tags:"Position" method:"get" summary:"get position list"`
	Sort   int `json:"sort" in:"query" dc:"sort"`
	CommonPaginationReq
}

type PositionGetListCommonRes struct {
	List  interface{} `json:"list" description:"position list"`
	Page  int         `json:"page" description:"page number"`
	Size  int         `json:"size" description:"size of page"`
	Total int         `json:"total" description:"total amount"`
}
