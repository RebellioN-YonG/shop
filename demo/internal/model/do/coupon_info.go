// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponInfo is the golang structure of table coupon_info for DAO operations like Where/Data.
type CouponInfo struct {
	g.Meta     `orm:"table:coupon_info, do:true"`
	Id         any         //
	Name       any         //
	Price      any         // 优惠前面值 单位分
	GoodsIds   any         // 关联使用的goods_ids  逗号分隔
	CategoryId any         // 关联使用的分类id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
