// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillOrder is the golang structure of table seckill_order for DAO operations like Where/Data.
type SeckillOrder struct {
	g.Meta         `orm:"table:seckill_order, do:true"`
	Id             any         // 主键ID
	OrderId        any         // 订单ID
	UserId         any         // 用户ID
	GoodsId        any         // 商品ID
	GoodsOptionsId any         // 商品规格ID
	SeckillPrice   any         // 秒杀价格 单位分
	Status         any         // 状态：0-待支付 1-已支付 2-已取消
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
