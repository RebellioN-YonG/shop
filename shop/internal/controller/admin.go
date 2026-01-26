package controller

import (
	"context"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"

	g "github.com/gogf/gf/v2/frame/g"
)

var Admin = cAdmin{}

type cAdmin struct{}

func (c *cAdmin) Create(ctx context.Context, req *backend.AdminReq) (res *backend.AdminRes, err error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminRes{
		AdminId: out.AdminId,
	}, nil
}

func (c *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	g.Log().Info(ctx, "Admin Update req: AdminId=%d, PicUrl=%s, Link=%s, Sort=%d", req.AdminId, req.PicUrl, req.Link, req.Sort)
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
		Id: req.AdminId,
	})
	if err != nil {
		g.Log().Error(ctx, "Admin Update err: %v", err)
		return nil, err
	}
	return &backend.AdminUpdateRes{Id: req.AdminId}, nil
}

func (c *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.AdminId)
	return
}

func (c *cAdmin) List(ctx context.Context, req *backend.AdminGetListCommonReq) (res *backend.AdminGetListCommonRes, err error) {
	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminGetListCommonRes{
		List:  getListRes.List,
		Total: getListRes.Total,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
	}, nil
}
