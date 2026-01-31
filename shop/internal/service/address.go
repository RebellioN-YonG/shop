// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/internal/model"
)

type (
	IAddress interface {
		// Add 添加地址
		Add(ctx context.Context, in *model.AddAddressInput) (out *model.AddAddressOutput, err error)
		// Update 更新地址
		Update(ctx context.Context, in *model.UpdateAddressInput) (err error)
		// Delete 删除地址
		Delete(ctx context.Context, id int) (err error)
		// Page 分页查询地址
		Page(ctx context.Context, in *model.PageAddressInput) (out *model.PageAddressOutput, err error)
		// GetCityList 获取城市地址列表（不分页）
		GetCityList(ctx context.Context) (out *model.CityAddressListOutput, err error)
	}
)

var (
	localAddress IAddress
)

func Address() IAddress {
	if localAddress == nil {
		panic("implement not found for interface IAddress, forgot register?")
	}
	return localAddress
}

func RegisterAddress(i IAddress) {
	localAddress = i
}
