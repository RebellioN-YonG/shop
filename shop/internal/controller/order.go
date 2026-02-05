package controller

var Order = cOrder{}

type cOrder struct{}

// todo: frontend
// func (c *cOrder) Add(ctx context.Context, req *end.OrderAddReq) (res *backend.OrderAddRes, err error) {

// 	out, err := service.Order().Create(ctx, &model.OrderCreateInput{
// 		OrderCreateUpdateBase: model.OrderCreateUpdateBase{
// 			Name: req.Name,
// 			Path: req.Path,
// 		},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &backend.OrderRes{OrderId: out.OrderId}, nil
// }

// func (c *cOrder) Delete(ctx context.Context, req *backend.OrderDeleteReq) (res *backend.OrderDeleteRes, err error) {
// 	err = service.Order().Delete(ctx, req.Id)
// 	return
// }

// func (c *cOrder) Update(ctx context.Context, req *backend.OrderUpdateReq) (res *backend.OrderUpdateRes, err error) {
// 	err = service.Order().Update(ctx, model.OrderUpdateInput{
// 		Id: req.Id,
// 		OrderCreateUpdateBase: model.OrderCreateUpdateBase{
// 			Name: req.Name,
// 			Path: req.Path,
// 		},
// 	})
// 	return &backend.OrderUpdateRes{Id: req.Id}, nil
// }

// func (c *cOrder) List(ctx context.Context, req *backend.OrderGetListCommonReq) (res *backend.OrderGetListCommonRes, err error) {
// 	getListRes, err := service.Order().GetList(ctx, &model.OrderGetListInput{
// 		Page: req.Page,
// 		Size: req.Size,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &backend.OrderGetListCommonRes{
// 		List:  getListRes.List,
// 		Page:  getListRes.Page,
// 		Size:  getListRes.Size,
// 		Total: getListRes.Total,
// 	}, nil
// }
