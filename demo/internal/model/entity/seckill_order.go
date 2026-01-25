// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillOrder is the golang structure for table seckill_order.
type SeckillOrder struct {
	Id             int64       `json:"id"             orm:"id"               description:"主键ID"`
	OrderId        int64       `json:"orderId"        orm:"order_id"         description:"订单ID"`
	UserId         int64       `json:"userId"         orm:"user_id"          description:"用户ID"`
	GoodsId        int64       `json:"goodsId"        orm:"goods_id"         description:"商品ID"`
	GoodsOptionsId int64       `json:"goodsOptionsId" orm:"goods_options_id" description:"商品规格ID"`
	SeckillPrice   int         `json:"seckillPrice"   orm:"seckill_price"    description:"秒杀价格 单位分"`
	Status         int         `json:"status"         orm:"status"           description:"状态：0-待支付 1-已支付 2-已取消"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`
}
