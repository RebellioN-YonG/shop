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
	IPermission interface {
		// Create 创建角色
		Create(ctx context.Context, in *model.PermissionCreateInput) (out model.PermissionCreateOutput, err error)
		// Update 更新角色
		Update(ctx context.Context, in model.PermissionUpdateInput) error
		// Delete 删除角色
		Delete(ctx context.Context, id uint) error
		// GetList 获取角色列表（含权限）
		GetList(ctx context.Context, in *model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error)
		// GetPathsByRoleIds 根据角色ID获取权限列表
		GetPathsByRoleIds(ctx context.Context, RoleIds []int) ([]string, error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
