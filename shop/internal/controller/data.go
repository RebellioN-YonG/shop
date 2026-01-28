package controller

import (
	"context"
	"shop/api/backend"
	"shop/internal/service"
)

var Data = cData{}

type cData struct{}

func (c *cData) HeadCard(ctx context.Context, req *backend.DataHeadReq) (res *backend.DataHeadRes, err error) {
	card, err := service.Data().HeadCard(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DataHeadRes{
		TodayOrderCount: card.TodayOrderCount,
		DAU:             card.DAU,
		ConversionRate:  card.ConversionRate,
	}, nil
}

func (c *cData) ECharts(ctx context.Context, req *backend.DataEChartsReq) (out *backend.DataEChartsRes, err error) {
	echarts, err := service.Data().ECharts(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DataEChartsRes{
		OrderTotal:           echarts.OrderTotal,
		SalePriceTotal:       echarts.SalePriceTotal,
		ConsumptionPerPerson: echarts.ConsumptionPerPerson,
		NewOrder:             echarts.NewOrder,
	}, nil
}
