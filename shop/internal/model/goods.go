package model

import (
	"shop/internal/model/do"
	"shop/internal/model/entity"
)

type GoodsCreateUpdateBase struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	Stock            int
	Sale             uint
	Tags             string
	DetailInfo       string
}

type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

// create goods output
type GoodsCreateOutput struct {
	Id uint `json:"id"`
}

type GoodsUpdateInput struct {
	GoodsCreateUpdateBase
	Id uint
}

type GoodsGetListInput struct {
	Page int
	Size int
	Sort int
}

type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem `json:"list" description:"list"`
	Total int                      `json:"total" description:"total"`
	Page  int                      `json:"page" description:"page"`
	Size  int                      `json:"size" description:"size"`
}

type GoodsGetListOutputItem struct {
	entity.GoodsInfo
}

type GoodsDetailInput struct {
	do.GoodsInfo
	IsCollect bool                  `json:"is_collect" description:"is collect"`
	Options   []do.GoodsOptionsInfo `json:"options" description:"options"`
	// todo:
	// Comments  []*CommentBase        `json:"comments" description:"comments"`
}

type GoodsDetailOutput struct {
	entity.GoodsInfo
}
