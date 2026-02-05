package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type OrderAddReq struct {
	g.Meta    `path:"/order/add" tags:"Order" method:"post" summary:"Add order"`
	GoodsList []OrderGoodsInfo `json:"goods_list" description:"Goods List"`
}

type OrderAddRes struct {
	Id uint `json:"id"`
}

type OrderDeleteReq struct {
	g.Meta `path:"/order/delete" tags:"Order" method:"post" summary:"Delete order"`
	Id     uint `json:"id" v:"required" description:"Order ID"`
}

type OrderDeleteRes struct {
	Id uint `json:"id"`
}

type OrderUpdateStatusReq struct {
	g.Meta `path:"/order/status" tags:"Order" method:"post" summary:"Update order status"`
	Id     uint `json:"id" v:"required" description:"Order ID"`
	Status int  `json:"status" v:"required" description:"Status"`
}

type OrderUpdateStatusRes struct {
	Id uint `json:"id"`
}

type OrderListReq struct {
	g.Meta `path:"/order/list" tags:"Order" method:"get" summary:"Get order list"`
	CommonPaginationReq
	Number         string      `json:"number" description:"Order Number"`
	UserId         int         `json:"user_id" description:"User ID"`
	PayType      int         `json:"pay_type" description:"Pay Type"`
	PayAtGte       *gtime.Time `json:"pay_at_gte" description:"Pay Time Start"`
	PayAtLte       *gtime.Time `json:"pay_at_lte" description:"Pay Time End"`
	Status         int         `json:"status" description:"Status"`
	ConsigneePhone string      `json:"consignee_phone" description:"Consignee Phone"`
	PriceGte       int         `json:"price_gte" description:"Price Start"`
	PriceLte       int         `json:"price_lte" description:"Price End"`
	DateGte        *gtime.Time `json:"dateGte" description:"Date Start"`
	DateLte        *gtime.Time `json:"dateLte" description:"Date End"`
}

type OrderListRes struct {
	CommonPaginationRes
}

type OrderDetailReq struct {
	g.Meta `path:"/order/detail" tags:"Order" method:"get" summary:"Get order detail"`
	Id     int `json:"id" v:"required" description:"order ID"`
}

type OrderBase struct {
	Id          int         `json:"id"               description:"order id"`
	Price       int         `json:"price"            description:"order amount"`
	CouponPrice int         `json:"coupon_price"      description:"coupon amount"`
	ActualPrice int         `json:"actual_price"      description:"actual payment amount"`
	Remark      string      `json:"remark"           description:"remarks"`
	CreatedAt   *gtime.Time `json:"created_at"        description:"create time"`
	UpdatedAt   *gtime.Time `json:"updated_at"        description:"update time"`
}

type OrderGoodsInfoBase struct {
	Id             int         `json:"id"             description:"order goods id"`
	OrderId        int         `json:"order_id"        description:"order id"`
	GoodsId        int         `json:"goods_id"        description:"goods id"`
	GoodsOptionsId int         `json:"goods_options_id" description:"goods options id sku id"`
	Count          int         `json:"count"          description:"goods count"`
	Remark         string      `json:"remark"         description:"remarks"`
	Price          int         `json:"price"          description:"order amount "`
	CouponPrice    int         `json:"coupon_price"    description:"coupon amount "`
	ActualPrice    int         `json:"actual_price"    description:"actual payment amount "`
	CreatedAt      *gtime.Time `json:"created_at"      description:"created at"`
	UpdatedAt      *gtime.Time `json:"updated_at"      description:"updated at"`
}

type OrderInfo struct {
	OrderBase
	Number           string      `json:"number"           description:"order number"`
	UserId           int         `json:"user_id"           description:"user id"`
	PayType        int         `json:"pay_type"          description:"payment method 1 wechat 2 alipay 3 yunshafu"`
	Status           int         `json:"status"           description:"order status: 1 pending payment 2 paid pending delivery 3 delivered 4 received pending evaluation 5 evaluated"`
	PayAt            *gtime.Time `json:"pay_at"            description:"payment time"`
	ShipAt           *gtime.Time `json:"ship_at"           description:"shipping time"`
	FinishAt         *gtime.Time `json:"finish_at"         description:"finish time"`
	ConsigneeName    string      `json:"consignee_name"    description:"consignee name"`
	ConsigneePhone   string      `json:"consignee_phone"   description:"consignee phone"`
	ConsigneeAddress string      `json:"consignee_address" description:"consignee address"`
}

type OrderGoodsInfo struct {
	GoodsId        int    `json:"goods_id" description:"goods id"`
	GoodsOptionsId int    `json:"goods_options_id" description:"goods options id sku id"`
	Count          int    `json:"count" description:"goods count"`
	Price          int    `json:"price" description:"order amount "`
	CouponPrice    int    `json:"coupon_price" description:"coupon amount "`
	ActualPrice    int    `json:"actual_price" description:"actual payment amount "`
	Remark         string `json:"remark" description:"remarks"`
}

type OrderRefundReq struct {
	g.Meta `path:"/order/refund" tags:"Order" method:"post" summary:"Order refund"`
	Id     uint   `json:"order_id" v:"required" description:"Order ID"`
	Reason string `json:"reason" v:"required" description:"Reason"`
}

type OrderRefundRes struct {
	Id uint `json:"id"`
}
