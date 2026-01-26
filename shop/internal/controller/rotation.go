package controller

import (
	"context"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"

	g "github.com/gogf/gf/v2/frame/g"
)

var Rotation = cRotation{}

type cRotation struct{}

func (c *cRotation) Create(ctx context.Context, req *backend.RotationReq) (res *backend.RotationRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationRes{
		RotationId: out.RotationId,
	}, nil
}

func (c *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	g.Log().Info(ctx, "Rotation Update req: RotationId=%d, PicUrl=%s, Link=%s, Sort=%d", req.RotationId, req.PicUrl, req.Link, req.Sort)
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
		Id: req.RotationId,
	})
	if err != nil {
		g.Log().Error(ctx, "Rotation Update err: %v", err)
		return nil, err
	}
	return &backend.RotationUpdateRes{Id: req.RotationId}, nil
}

func (c *cRotation) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.RotationId)
	return
}

func (c *cRotation) List(ctx context.Context, req *backend.RotationGetListCommonReq) (res *backend.RotationGetListCommonRes, err error) {
	getListRes, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationGetListCommonRes{
		List:  getListRes.List,
		Total: getListRes.Total,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
	}, nil
}
