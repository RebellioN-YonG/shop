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
	IRole interface {
		// Create 创建角色
		Create(ctx context.Context, in *model.RoleCreateInput) (out model.RoleCreateOutput, err error)
		// Update 更新角色
		Update(ctx context.Context, in model.RoleUpdateInput) error
		// Delete 删除角色
		Delete(ctx context.Context, id uint) error
		// GetList 获取角色列表（含权限）
		GetList(ctx context.Context, in *model.RoleGetListInput) (out *model.RoleGetListOutput, err error)
		// GetPermissionsByRoleId 根据角色ID获取权限列表
		GetPermissionsByRoleId(ctx context.Context, roleId uint) ([]model.RolePermissionItem, error)
		// AddPermission 为角色添加单个权限
		AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error)
		// AddPermissions 批量为角色添加权限
		AddPermissions(ctx context.Context, roleId uint, permissionIds []uint) error
		// DeletePermission 删除角色权限
		DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error
		DeletePermissions(ctx context.Context, roleId uint, permissionIds []uint) error
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
