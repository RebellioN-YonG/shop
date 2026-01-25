// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsOptionsInfo is the golang structure of table goods_options_info for DAO operations like Where/Data.
type GoodsOptionsInfo struct {
	g.Meta    `orm:"table:goods_options_info, do:true"`
	Id        any         //
	GoodsId   any         // 商品id
	PicUrl    any         // 图片
	Name      any         // 商品名称
	Price     any         // 价格 单位分
	Stock     any         // 库存
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
