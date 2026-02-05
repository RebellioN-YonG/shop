package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type OrderInfo struct {
	Id               uint        `json:"id"             description:"order goods id"`
	UserId           uint        `json:"user_id"           description:"user id"`
	Number           string      `json:"number"           description:"order number"`
	Remark           string      `json:"remark"         description:"remarks"`
	Price            int         `json:"price"          description:"order amount "`
	CouponPrice      int         `json:"coupon_price"    description:"coupon amount "`
	ActualPrice      int         `json:"actual_price"    description:"actual payment amount "`
	PayType          int         `json:"pay_type"          description:"payment method 1 wechat 2 alipay 3 yunshafu"`
	Status           int         `json:"status"           description:"order status: 1 pending payment 2 paid pending delivery 3 delivered 4 received pending evaluation 5 evaluated"`
	CreatedAt        *gtime.Time `json:"created_at"      description:"created at"`
	UpdatedAt        *gtime.Time `json:"updated_at"      description:"updated at"`
	PayAt            *gtime.Time `json:"pay_at"            description:"payment time"`
	ConsigneeName    string      `json:"consignee_name"    description:"consignee name"`
	ConsigneePhone   string      `json:"consignee_phone"   description:"consignee phone"`
	ConsigneeAddress string      `json:"consignee_address" description:"consignee address"`
}

type OrderGoodsInfo struct {
	Id             int    `json:"id" description:"order goods id"`
	OrderId        int    `json:"order_id" description:"order id"`
	GoodsId        int    `json:"goods_id" description:"goods id"`
	GoodsOptionsId int    `json:"goods_options_id" description:"goods options id sku id"`
	Count          int    `json:"count" description:"goods count"`
	Price          int    `json:"price" description:"order amount "`
	CouponPrice    int    `json:"coupon_price" description:"coupon amount "`
	ActualPrice    int    `json:"actual_price" description:"actual payment amount "`
	Remark         string `json:"remark" description:"remarks"`
}

type OrderAddGoodsInfo struct {
	GoodsId        int    `json:"goods_id"`
	GoodsOptionsId int    `json:"goods_options_id"`
	Count          int    `json:"count"`
	Remark         string `json:"remark"`
	Price          int    `json:"price"`
	CouponPrice    int    `json:"coupon_price"`
	ActualPrice    int    `json:"actual_price"`
}

type OrderAddReq struct {
	g.Meta           `path:"/order/add" tags:"frontend order" method:"post" summary:"Add order"`
	Price            int    `json:"price" description:"Price"`
	CouponPrice      int    `json:"coupon_price" description:"Coupon Price"`
	ActualPrice      int    `json:"actual_price" description:"Actual Price"`
	Remark           string `json:"remark" description:"Remark"`
	ConsigneeName    string `json:"consignee_name" description:"Consignee Name"`
	ConsigneePhone   string `json:"consignee_phone" description:"Consignee Phone"`
	ConsigneeAddress string `json:"consignee_address" description:"Consignee Address"`
	AddressId        uint   `json:"address_id" description:"Address ID"`

	OrderAddGoodsInfos []*OrderAddGoodsInfo `json:"order_add_goods_infos"`
	GoodsList          []OrderAddGoodsInfo  `json:"goods_list" description:"Goods list"`
}

type OrderAddRes struct {
	Id uint `json:"id"`
}

type OrderListReq struct {
	g.Meta `path:"/order/list" tags:"frontend order" method:"get" summary:"Get order list"`
	CommonPaginationReq
	Status  int    `json:"status" description:"Status"`
	Page    int    `json:"page" description:"Page"`
	Size    int    `json:"size" description:"Size"`
	DateGte string `json:"date_gte" description:"Date Start"`
	DateLte string `json:"date_lte" description:"Date End"`
}

type OrderListRes struct {
	CommonPaginationRes
	List []OrderInfo `json:"list" description:"Order list"`
}

type OrderDetailReq struct {
	g.Meta `path:"/order/detail" tags:"Order" method:"get" summary:"Get order detail"`
	Id     int `json:"id" v:"required" description:"order ID"`
}

type OrderDetailRes struct {
	OrderInfo OrderInfo        `json:"order_info" description:"Order info"`
	GoodsInfo []OrderGoodsInfo `json:"goods_info" description:"Goods info"`
}

type OrderRefundReq struct {
	g.Meta  `path:"/order/refund" tags:"Order" method:"post" summary:"Order refund"`
	OrderId int    `json:"order_id" v:"required" description:"Order ID"`
	Reason  string `json:"reason" v:"required" description:"Reason"`
}

type OrderRefundRes struct {
	Id uint `json:"id"`
}

type OrderCancelReq struct {
	g.Meta `path:"/order/cancel" tags:"Order" method:"post" summary:"Order cancel"`
	Id     uint   `json:"id" v:"required" description:"Order ID"`
	Reason string `json:"reason" v:"required" description:"Reason"`
}

type OrderCancelRes struct {
	Id uint `json:"id"`
}

type OrderConfirmReq struct {
	g.Meta `path:"/order/confirm" tags:"Order" method:"post" summary:"Order confirm"`
	Id     uint `json:"id" v:"required" description:"Order ID"`
}

type OrderConfirmRes struct {
	Id uint `json:"id"`
}
type OrderPayReq struct {
	g.Meta  `path:"/order/pay" tags:"Order" method:"post" summary:"Order pay"`
	Id      uint `json:"id" v:"required" description:"Order ID"`
	PayType int  `json:"pay_type" v:"required" description:"Pay Type: 1 wechat 2 alipay 3 unionpay"`
}

type OrderPayRes struct {
	PayUrl string `json:"pay_url" description:"Pay URL"`
}
