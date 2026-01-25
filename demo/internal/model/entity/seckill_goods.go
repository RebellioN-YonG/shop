// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillGoods is the golang structure for table seckill_goods.
type SeckillGoods struct {
	Id             int64       `json:"id"             orm:"id"               description:"主键ID"`
	GoodsId        int64       `json:"goodsId"        orm:"goods_id"         description:"商品ID"`
	GoodsOptionsId int64       `json:"goodsOptionsId" orm:"goods_options_id" description:"商品规格ID"`
	OriginalPrice  int         `json:"originalPrice"  orm:"original_price"   description:"原始价格 单位分"`
	SeckillPrice   int         `json:"seckillPrice"   orm:"seckill_price"    description:"秒杀价格 单位分"`
	SeckillStock   int         `json:"seckillStock"   orm:"seckill_stock"    description:"秒杀库存"`
	StartTime      *gtime.Time `json:"startTime"      orm:"start_time"       description:"秒杀开始时间"`
	EndTime        *gtime.Time `json:"endTime"        orm:"end_time"         description:"秒杀结束时间"`
	Status         int         `json:"status"         orm:"status"           description:"状态：0-未开始 1-进行中 2-已结束"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`
}
