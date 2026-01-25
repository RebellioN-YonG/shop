// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillGoods is the golang structure of table seckill_goods for DAO operations like Where/Data.
type SeckillGoods struct {
	g.Meta         `orm:"table:seckill_goods, do:true"`
	Id             any         // 主键ID
	GoodsId        any         // 商品ID
	GoodsOptionsId any         // 商品规格ID
	OriginalPrice  any         // 原始价格 单位分
	SeckillPrice   any         // 秒杀价格 单位分
	SeckillStock   any         // 秒杀库存
	StartTime      *gtime.Time // 秒杀开始时间
	EndTime        *gtime.Time // 秒杀结束时间
	Status         any         // 状态：0-未开始 1-进行中 2-已结束
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
