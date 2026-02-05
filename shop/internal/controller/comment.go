package controller

// todo:

// import (
// 	"context"

// 	"shop/api/backend"
// 	"shop/internal/model"
// 	"shop/internal/service"
// )

var Comment = cComment{}

type cComment struct{}

// func (c *cComment) Create(ctx context.Context, req *backend.CommentReq) (res *backend.CommentRes, err error) {
// 	out, err := service.Comment().Create(ctx, &model.CommentCreateInput{
// 		CommentCreateUpdateBase: model.CommentCreateUpdateBase{
// 			Name: req.Name,
// 			Path: req.Path,
// 		},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &backend.CommentRes{CommentId: out.CommentId}, nil
// }

// func (c *cComment) Delete(ctx context.Context, req *backend.CommentDeleteReq) (res *backend.CommentDeleteRes, err error) {
// 	err = service.Comment().Delete(ctx, req.Id)
// 	return
// }

// func (c *cComment) Update(ctx context.Context, req *backend.CommentUpdateReq) (res *backend.CommentUpdateRes, err error) {
// 	err = service.Comment().Update(ctx, model.CommentUpdateInput{
// 		Id: req.Id,
// 		CommentCreateUpdateBase: model.CommentCreateUpdateBase{
// 			Name: req.Name,
// 			Path: req.Path,
// 		},
// 	})
// 	return &backend.CommentUpdateRes{Id: req.Id}, nil
// }

// func (c *cComment) List(ctx context.Context, req *backend.CommentGetListCommonReq) (res *backend.CommentGetListCommonRes, err error) {
// 	getListRes, err := service.Comment().GetList(ctx, &model.CommentGetListInput{
// 		Page: req.Page,
// 		Size: req.Size,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &backend.CommentGetListCommonRes{
// 		List:  getListRes.List,
// 		Page:  getListRes.Page,
// 		Size:  getListRes.Size,
// 		Total: getListRes.Total,
// 	}, nil
// }
