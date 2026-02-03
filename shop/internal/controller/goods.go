package controller

// todo
// import (
// 	"context"

// 	"shop/api/backend"
// 	"shop/internal/model"
// 	"shop/internal/service"
// )

var Goods = cGoods{}

type cGoods struct{}

// func (c *cGoods) Create(ctx context.Context, req *backend.GoodsReq) (res *backend.GoodsRes, err error) {
// 	out, err := service.Goods().Create(ctx, &model.GoodsCreateInput{
// 		GoodsCreateUpdateBase: model.GoodsCreateUpdateBase{
// 			Name: req.Name,
// 			Path: req.Path,
// 		},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &backend.GoodsRes{GoodsId: out.GoodsId}, nil
// }

// func (c *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
// 	err = service.Goods().Delete(ctx, req.Id)
// 	return
// }

// func (c *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
// 	err = service.Goods().Update(ctx, model.GoodsUpdateInput{
// 		Id: req.Id,
// 		GoodsCreateUpdateBase: model.GoodsCreateUpdateBase{
// 			Name: req.Name,
// 			Path: req.Path,
// 		},
// 	})
// 	return &backend.GoodsUpdateRes{Id: req.Id}, nil
// }

// func (c *cGoods) List(ctx context.Context, req *backend.GoodsGetListCommonReq) (res *backend.GoodsGetListCommonRes, err error) {
// 	getListRes, err := service.Goods().GetList(ctx, &model.GoodsGetListInput{
// 		Page: req.Page,
// 		Size: req.Size,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &backend.GoodsGetListCommonRes{
// 		List:  getListRes.List,
// 		Page:  getListRes.Page,
// 		Size:  getListRes.Size,
// 		Total: getListRes.Total,
// 	}, nil
// }
