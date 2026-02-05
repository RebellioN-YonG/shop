package model

import (
	"shop/internal/model/do"
	"shop/internal/model/entity"
)

// OrderListInput 获取订单列表输入
type OrderListInput struct {
	Page           int    `json:"page" description:"页码"`
	Size           int    `json:"size" description:"每页数量"`
	Number         string `json:"number" description:"订单编号"`
	UserId         int    `json:"user_id" description:"用户id"`
	PayType      int    `json:"pay_type" description:"支付方式"`
	Status         int    `json:"status" description:"订单状态"`
	PriceGte       int    `json:"price_gte" description:"金额开始"`
	PriceLte       int    `json:"price_lte" description:"金额结束"`
	PayAtLte       string `json:"pay_at_lte" description:"支付时间结束"`
	PayAtGte       string `json:"pay_at_gte" description:"支付时间开始"`
	DateGte        string `json:"date_gte" description:"创建时间开始"`
	DateLte        string `json:"date_lte" description:"创建时间结束"`
	ConsigneePhone string `json:"consignee_phone" description:"收货人手机号"`
}

// OrderListOutput 获取订单列表输出
type OrderListOutput struct {
	List  []OrderListOutputItem `json:"list" description:"列表"`
	Total int                   `json:"total" description:"总数"`
	Page  int                   `json:"page" description:"页码"`
	Size  int                   `json:"size" description:"每页数量"`
}

// OrderListOutputItem 订单列表项
type OrderListOutputItem struct {
	entity.OrderInfo
}

// OrderAddInput 创建订单输入
type OrderAddInput struct {
	UserId            uint   `json:"userId"`
	Number            string `json:"number"`
	Remark            string `json:"remark"`
	Price             int    `json:"price"`
	CouponPrice       int    `json:"coupon_price"`
	ActualPrice       int    `json:"actual_price"`
	ConsigneeName     string `json:"consignee_name"`
	ConsigneePhone    string `json:"consignee_phone"`
	ConsigneeAddress  string `json:"consignee_address"`
	OrderAddGoodsInfos []*OrderAddGoodsInfo
}

type OrderAddGoodsInfo struct {
	Id             int
	OrderId        int
	GoodsId        int
	GoodsOptionsId int
	Count          int
	Price          int
	CouponPrice    int
	ActualPrice    int
	Remark         string
}

// OrderAddOutput 创建订单输出
type OrderAddOutput struct {
	Id uint `json:"id" description:"订单id"`
}

// OrderDetailInput 获取订单详情输入
type OrderDetailInput struct {
	Id uint `json:"id" description:"订单id"`
}

// OrderDetailOutput 获取订单详情输出
type OrderDetailOutput struct {
	do.OrderInfo
	GoodsInfo []*do.OrderGoodsInfo `orm:"with:order_id=id"`
}

// SeckillOrderInput 秒杀订单输入
type SeckillOrderInput struct {
	UserId         uint    `json:"user_id"`
	GoodsId        uint    `json:"goods_id"`
	GoodsOptionsId uint    `json:"goods_options_id"`
	Count          uint    `json:"count"`
	Price          float64 `json:"price"`
	OrderNo        string  `json:"order_no"`
}

type OrderCreateInput struct {
	UserId         uint    `json:"user_id"`
	GoodsId        uint    `json:"goods_id"`
	GoodsOptionsId uint    `json:"goods_options_id"`
	Count          uint    `json:"count"`
	Price          float64 `json:"price"`
	Status         int     `json:"status"`
	OrderNo        string  `json:"order_no"`
	Address        string  `json:"address"`
	Phone          string  `json:"phone"`
	Remark         string  `json:"remark"`
}
