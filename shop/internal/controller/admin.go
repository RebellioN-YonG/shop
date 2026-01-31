package controller

import (
	"context"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

var Admin = cAdmin{}

type cAdmin struct{}

func (c *cAdmin) Create(ctx context.Context, req *backend.AdminReq) (res *backend.AdminRes, err error) {
	input := &model.AdminCreateInput{}
	input.Name = req.Name
	input.Password = req.Password
	input.RoleIds = req.RoleIds
	input.IsAdmin = req.IsAdmin
	out, err := service.Admin().Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return &backend.AdminRes{
		AdminId: out.AdminId,
	}, nil
}

func (c *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, &model.AdminUpdateInput{
		Id: req.AdminId,
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	return &backend.AdminUpdateRes{Id: req.AdminId}, nil
}

func (c *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.AdminId)
	return
}

func (c *cAdmin) List(ctx context.Context, req *backend.AdminGetListCommonReq) (res *backend.AdminGetListCommonRes, err error) {
	getListRes, err := service.Admin().GetList(ctx, &model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
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

func (c *cAdmin) Info(ctx context.Context, req *backend.AdminGetInfoReq) (res *backend.AdminGetInfoRes, err error) {
	return &backend.AdminGetInfoRes{
		Id:      gconv.Int(service.Session().GetUser(ctx).Id),
		IdKey:   service.Auth().IdentityKey,
		PayLoad: service.Auth().GetPayload(ctx),
	}, nil
}
