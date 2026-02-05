// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/api/backend"
	"shop/internal/model"

	"github.com/gogf/gf/v2/container/gmap"
)

type (
	IOrder interface {
		Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, in *backend.OrderDeleteReq) (out *backend.OrderDeleteRes, err error)
		// List: 获取订单列表
		List(ctx context.Context, in model.OrderListInput) (out *model.OrderListOutput, err error)
		OrderListCondition(ctx context.Context, in model.OrderListInput) *gmap.Map
		// 商品详情
		Detail(ctx context.Context, in model.OrderDetailInput) (out model.OrderDetailOutput, err error)
		// Refund: 退款
		Refund(ctx context.Context, in *backend.OrderRefundReq) (out *backend.OrderRefundRes, err error)
		// CreateSeckillOrder: 创建秒杀订单
		CreateSeckillOrder(ctx context.Context, in model.SeckillOrderInput) (out *model.OrderAddOutput, err error)
		// RestoreGoodsStock: 恢复商品库存
		RestoreGoodsStock(ctx context.Context, orderId uint) (err error)
		// GetOrderById: 根据订单ID获取订单详情
		GetOrderById(ctx context.Context, id uint) (order *model.OrderDetailOutput, err error)
		// UpdateOrderStatus: 更新订单状态
		UpdateOrderStatus(ctx context.Context, in *backend.OrderUpdateStatusReq) (out *backend.OrderUpdateStatusRes, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
