package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsCommonAddUpdate struct {
	PicUrl           string `json:"pic_url"`
	Name             string `json:"name" v:"required#goods name cannot be empty"`
	Price            int    `json:"price" v:"required#goods price cannot be empty"`
	Level1CategoryId uint   `json:"level_1_category_id"`
	Level2CategoryId uint   `json:"level_2_category_id"`
	Level3CategoryId uint   `json:"level_3_category_id"`
	Brand            string `json:"brand" v:"max-length:20#brand name cannot exceed 20 characters"`
	Stock            int    `json:"stock"`
	Sale             uint   `json:"sale"`
	Tags             string `json:"tags"`
	DetailInfo       string `json:"detail_info"`
}

type GoodsReq struct {
	g.Meta `path:"/goods/add" tags:"Goods" method:"post" summary:"ur 1st goods api"`
	GoodsCommonAddUpdate
}
type GoodsRes struct {
	Id uint `json:"id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" tags:"Goods" method:"delete" summary:"delete a goods"`
	Id     uint `v:"min:1#select a goods to delete" dc:"goods's id"`
}

type GoodsDeleteRes struct {
}

type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update" tags:"Goods" method:"post" summary:"update a goods"`
	Id     uint `json:"goods_id" v:"min:1#select a goods to update" dc:"goods's id"`
	GoodsCommonAddUpdate
}

type GoodsUpdateRes struct {
	Id uint `json:"id"`
}

type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" tags:"Goods" method:"get" summary:"get goods list"`
	CommonPaginationReq
}

type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"goods list"`
	Page  int         `json:"page" description:"page number"`
	Size  int         `json:"size" description:"size of page"`
	Total int         `json:"total" description:"total amount"`
}
