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
	IUser interface {
		Register(ctx context.Context, in *model.UserRegisterInput) (out model.UserRegisterOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// UpdateStatus 修改状态
		UpdateStatus(ctx context.Context, id uint, status int) (err error)
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.UserGetListInput) (out *model.UserGetListOutput, err error)
		// GetById 根据ID获取管理员信息
		GetById(ctx context.Context, id int) (*model.UserInfoItem, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
