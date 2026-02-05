package order

import (
	"context"
	"shop/api/backend"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	utility "shop/utility/utils"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

func (s *sOrder) Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	in.Number = utility.GetOrderNum()
	out = &model.OrderAddOutput{}
	lastInsertId, err := dao.OrderInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, err
	}
	for _, info := range in.OrderAddGoodsInfos {
		info.OrderId = gconv.Int(lastInsertId)
		_, err := dao.OrderGoodsInfo.Ctx(ctx).Insert(info)
		if err != nil {
			return nil, err
		}
	}
	// todo: update goods' sales and inventory
	return
}

// Delete 删除
func (s *sOrder) Delete(ctx context.Context, in *backend.OrderDeleteReq) (out *backend.OrderDeleteRes, err error) {
	// check order is exist
	order, err := s.GetOrderById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, gerror.New("order not exist")
	}
	// process with transaction
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// delete order goods
		_, err := dao.OrderGoodsInfo.Ctx(ctx).
			Where(dao.OrderGoodsInfo.Columns().OrderId, in.Id).
			Delete()
		if err != nil {
			return err
		}
		// delete order
		_, err = dao.OrderInfo.Ctx(ctx).
			Where(dao.OrderInfo.Columns().Id, in.Id).
			Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &backend.OrderDeleteRes{
		Id: in.Id,
	}, nil
}

// List: 获取订单列表
func (s *sOrder) List(ctx context.Context, in model.OrderListInput) (out *model.OrderListOutput, err error) {

	//1.获得*gdb.Model对象，方面后续调用
	m := dao.OrderInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.OrderListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.OrderListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// OrderListCondition: 订单列表筛选条件

func (s *sOrder) OrderListCondition(ctx context.Context, in model.OrderListInput) *gmap.Map {
	m := gmap.New()

	if in.Number != "" {
		m.Set(dao.OrderInfo.Columns().Number, in.Number)
	}
	if in.UserId != 0 {
		m.Set(dao.OrderInfo.Columns().UserId, in.UserId)
	}
	if in.PayType != 0 {
		m.Set(dao.OrderInfo.Columns().PayType, in.PayType)
	}
	if in.Status != 0 {
		m.Set(dao.OrderInfo.Columns().Status, in.Status)
	}
	if in.ConsigneePhone != "" {
		m.Set(dao.OrderInfo.Columns().ConsigneePhone+" like ", "%"+in.ConsigneePhone+"%")
	}
	if in.PriceGte != 0 {
		m.Set(dao.OrderInfo.Columns().Price+" >= ", in.PriceGte)
	}
	if in.PriceLte != 0 {
		m.Set(dao.OrderInfo.Columns().Price+" <= ", in.PriceLte)
	}
	if in.DateGte != "" {
		m.Set(dao.OrderInfo.Columns().CreatedAt+" >= ", gtime.New(in.DateGte).StartOfDay())
	}
	if in.DateLte != "" {
		m.Set(dao.OrderInfo.Columns().CreatedAt+" <= ", gtime.New(in.DateLte).EndOfDay())
	}
	return m
}

// 商品详情
func (*sOrder) Detail(ctx context.Context, in model.OrderDetailInput) (out model.OrderDetailOutput, err error) {
	return
}

// todo: frontend 交互部分

// Pay: 支付
// func (s *sOrder) Pay(ctx context.Context, in *frontend.OrderPayReq) (out *frontend.OrderPayRes, err error) {
// 	return
// }

// Cancel: 取消订单
// func (s *sOrder) Cancel(ctx context.Context, in *frontend.OrderCancelReq) (out *frontend.OrderCancelRes, err error) {
// 	return
// }

// Confirm: 确认订单
// func (s *sOrder) Confirm(ctx context.Context, in *frontend.OrderConfirmReq) (out *frontend.OrderConfirmRes, err error) {
// 	return
// }

// Refund: 退款
func (s *sOrder) Refund(ctx context.Context, in *backend.OrderRefundReq) (out *backend.OrderRefundRes, err error) {
	// check order
	order, err := s.GetOrderById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, gerror.New("order not exist")
	}
	// process with transaction
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// delete order goods
		_, err := dao.OrderGoodsInfo.Ctx(ctx).
			Data(g.Map{
				dao.OrderInfo.Columns().Status: 6,
				dao.OrderInfo.Columns().Remark: in.Reason,
			}).Where(dao.OrderGoodsInfo.Columns().Id, in.Id).
			Update()
		if err != nil {
			return err
		}
		// get order goods
		var orderGoods []*entity.OrderGoodsInfo
		err = dao.OrderGoodsInfo.Ctx(ctx).
			Where(dao.OrderGoodsInfo.Columns().OrderId, in.Id).
			Scan(&orderGoods)
		if err != nil {
			return err
		}
		// restore goods stock
		for _, goods := range orderGoods {
			_, err := dao.GoodsInfo.Ctx(ctx).
				WherePri(goods.GoodsId).
				Increment(dao.GoodsInfo.Columns().Stock, goods.Count)
			if err != nil {
				return err
			}
			if goods.GoodsOptionsId > 0 {
				_, err := dao.GoodsOptionsInfo.Ctx(ctx).
					WherePri(goods.GoodsOptionsId).
					Increment(dao.GoodsOptionsInfo.Columns().Stock, goods.Count)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &backend.OrderRefundRes{
		Id: in.Id,
	}, nil
}

// CreateSeckillOrder: 创建秒杀订单
func (s *sOrder) CreateSeckillOrder(ctx context.Context, in model.SeckillOrderInput) (out *model.OrderAddOutput, err error) {
	g.Log().Info(ctx, "Start Creating Seckill Order, OrderNo:%s, UserId:%d, GoodsId:%d",
		in.OrderNo, in.UserId, in.GoodsId)

	// conv to standard order
	orderAddInput := model.OrderAddInput{
		UserId:      in.UserId,
		Number:      in.OrderNo,
		Remark:      "秒杀订单",
		Price:       int(in.Price * float64(in.Count) * 100),
		ActualPrice: int(in.Price * float64(in.Count) * 100),
		OrderAddGoodsInfos: []*model.OrderAddGoodsInfo{
			{
				GoodsId:        int(in.GoodsId),
				GoodsOptionsId: int(in.GoodsOptionsId),
				Count:          int(in.Count),
				Price:          int(in.Price * 100),
				ActualPrice:    int(in.Price * 100),
			},
		},
	}
	out, err = s.Add(ctx, orderAddInput)
	return
}

// ProcessOrderMessage: 处理订单消息
// func (s *sOrder) ProcessOrderMessage(ctx context.Context, msg []byte) error {
// 	// 解析订单消息
// 	// todo: 转换秒杀订单
// 	// var orderMsg model.SeckillOrderMsg
// 	// 普通订单输入转换秒杀输入

// 	// 创建订单
// 	return nil
// }

// RestoreGoodsStock: 恢复商品库存
func (s *sOrder) RestoreGoodsStock(ctx context.Context, orderId uint) (err error) {
	// 查询订单商品
	var orderGoods []*model.OrderAddGoodsInfo
	err = dao.OrderGoodsInfo.Ctx(ctx).
		Where(dao.OrderGoodsInfo.Columns().OrderId, orderId).
		Scan(&orderGoods)
	if err != nil {
		return err
	}
	// 使用事务处理
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, item := range orderGoods {
			// 恢复商品库存
			_, err := dao.GoodsInfo.Ctx(ctx).
				WherePri(item.GoodsId).
				Increment(dao.GoodsInfo.Columns().Stock, item.Count)
			if err != nil {
				return err
			}
			// 恢复商品规格库存
			_, err = dao.GoodsOptionsInfo.Ctx(ctx).
				WherePri(item.GoodsOptionsId).
				Increment(dao.GoodsOptionsInfo.Columns().Stock, item.Count)
			if err != nil {
				return err
			}
			// 恢复商品销量
			_, err = dao.GoodsInfo.Ctx(ctx).
				WherePri(item.GoodsId).
				Decrement(dao.GoodsInfo.Columns().Sale, item.Count)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

// GetOrderById: 根据订单ID获取订单详情
func (s *sOrder) GetOrderById(ctx context.Context, id uint) (order *model.OrderDetailOutput, err error) {
	err = dao.OrderInfo.Ctx(ctx).WherePri(id).Scan(&order)
	return
}

// UpdateOrderStatus: 更新订单状态
func (s *sOrder) UpdateOrderStatus(ctx context.Context, in *backend.OrderUpdateStatusReq) (out *backend.OrderUpdateStatusRes, err error) {
	return
}
