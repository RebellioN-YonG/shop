package controller

import (
	"context"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

var Address = cAddress{}

type cAddress struct{}

// Add 添加地址
func (c *cAddress) Add(ctx context.Context, req *backend.AddAddressReq) (res *backend.AddAddressRes, err error) {
	out, err := service.Address().Add(ctx, &model.AddAddressInput{
		AddressBase: model.AddressBase{
			ParentId: req.ParentId,
			Name:     req.Name,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AddAddressRes{Id: out.Id}, nil
}

// Update 更新地址
func (c *cAddress) Update(ctx context.Context, req *backend.UpdateAddressReq) (res *backend.UpdateAddressRes, err error) {
	err = service.Address().Update(ctx, &model.UpdateAddressInput{
		Id: req.Id,
		AddressBase: model.AddressBase{
			ParentId: req.ParentId,
			Name:     req.Name,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.UpdateAddressRes{Id: req.Id}, nil
}

// Delete 删除地址
func (c *cAddress) Delete(ctx context.Context, req *backend.DeleteAddressReq) (res *backend.DeleteAddressRes, err error) {
	err = service.Address().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.DeleteAddressRes{}, nil
}

// Page 分页查询地址
func (c *cAddress) Page(ctx context.Context, req *backend.PageAddressReq) (res *backend.PageAddressRes, err error) {
	out, err := service.Address().Page(ctx, &model.PageAddressInput{
		CommonPaginationReq: backend.CommonPaginationReq{
			Page: req.Page,
			Size: req.Size,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PageAddressRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  out.List,
			Total: out.Total,
			Page:  out.Page,
			Size:  out.Size,
		},
	}, nil
}

// GetCityList 获取城市地址列表
func (c *cAddress) GetCityList(ctx context.Context, req *backend.CityAddressReq) (res *backend.CityAddressRes, err error) {
	out, err := service.Address().GetCityList(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.CityAddressRes{List: out.List}, nil
}
