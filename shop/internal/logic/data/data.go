package data

import (
	"context"
	"fmt"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
	utility "shop/utility/utils"

	"github.com/gogf/gf/v2/os/gtime"
)

type sData struct{}

func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) HeadCard(ctx context.Context) (out *model.HeadDataOutput, err error) {
	return &model.HeadDataOutput{
		TodayOrderCount: TodayOrderCount(ctx),
		DAU:             utility.RandInt(200),
		ConversionRate:  utility.RandInt(80),
	}, nil
}

func (s *sData) ECharts(ctx context.Context) (out *model.EChartsOutput, err error) {
	return &model.EChartsOutput{

		OrderTotal:           OrderTotal(ctx),
		SalePriceTotal:       SalePriceTotalRecentDays(ctx),
		ConsumptionPerPerson: OrderTotal(ctx),
		NewOrder:             OrderTotal(ctx),
	}, nil
}

func TodayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt,
			gtime.Now().StartOfDay(),
			gtime.Now().EndOfDay()).Count("id")
	if err != nil {
		return -1
	}
	return
}

/*
  - SELECT COUNT(*) total,age FROM `user` GROUP BY age
    db.Model("user").
    Fields("COUNT(*) total,age").
    Group("age").All()
*/
func OrderTotal(ctx context.Context) (counts []int) {
	counts = make([]int, 0, 7)
	recent7Dates := utility.GetRecent7Date()
	var TodayTotals []model.TodayTotal
	err := dao.OrderInfo.Ctx(ctx).
		Where(dao.OrderInfo.Columns().CreatedAt+" >= ", utility.GetBefore7Date()).
		Fields("count(*) total, date_format(created_at, '%Y-%m-%d') today").
		Group("today").
		Scan(&TodayTotals)
	fmt.Printf("result:%v", TodayTotals)
	for i, date := range recent7Dates {
		for _, todayTotal := range TodayTotals {
			if date == todayTotal.Today {
				counts[i] = todayTotal.Total
			}
		}
	}
	if err != nil {
		return counts
	}
	return
}

func SalePriceTotalRecentDays(ctx context.Context) (totals []int) {
	totals = make([]int, 0, 7)
	recent7Dates := utility.GetRecent7Date()
	var TodayTotals []model.TodayTotal
	err := dao.OrderInfo.Ctx(ctx).
		Where(dao.OrderInfo.Columns().CreatedAt+" >= ", utility.GetBefore7Date()).
		Fields("sum(actual_price) total, date_format(created_at, '%Y-%m-%d') today").
		Group("today").
		Scan(&TodayTotals)
	fmt.Printf("result:%v", TodayTotals)
	for i, date := range recent7Dates {
		for _, todayTotal := range TodayTotals {
			if date == todayTotal.Today {
				totals[i] = todayTotal.Total
			}
		}
	}
	if err != nil {
		return totals
	}
	return
}
