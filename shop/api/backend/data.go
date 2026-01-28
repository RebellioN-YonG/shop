package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DataHeadReq struct {
	g.Meta `path:"/dashborad/head" tags:"Data" method:"post" summary:"execute login request"`
}

type DataHeadRes struct {
	TodayOrderCount int `json:"total_order_count" description:"today order count"`
	DAU             int `json:"dau" desc:"DAU"`
	ConversionRate  int `json:"conversion_rate" desc:"conversion rate"`
}

type DataEChartsReq struct {
	g.Meta `path:"/dashborad/echarts" method:"post" tags:"data" summary:"get data for echarts" sunmary:"get data for echarts"`
}

type DataEChartsRes struct {
	OrderTotal           []int `json:"order_total" desc:"order count"`
	SalePriceTotal       []int `json:"sale_price_total" desc:"sale price"`
	ConsumptionPerPerson []int `json:"consumption_per_person" desc:"avg consumption"`
	NewOrder             []int `json:"new_order" desc:"new order"`
}
