package model

type HeadDataOutput struct {
	TodayOrderCount int `json:"today_order_count" desc:"today order count"`
	DAU             int `json:"dau" desc:"DAU"`
	ConversionRate  int `json:"conversion_rate" desc:"conversion rate"`
}
type EChartsOutput struct {
	OrderTotal           []int `json:"order_total" desc:"order count"`
	SalePriceTotal       []int `json:"sale_price_total" desc:"sale price"`
	ConsumptionPerPerson []int `json:"consumption_per_person" desc:"avg consumption"`
	NewOrder             []int `json:"new_order" desc:"new order"`
}

type TodayTotal struct {
	Today string `json:"today"`
	Total int    `json:"total"`
}
