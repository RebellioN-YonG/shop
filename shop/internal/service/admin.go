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
	IAdmin interface {
		Create(ctx context.Context, in *model.AdminCreateInput) (out model.AdminCreateOutput, err error)
		Update(ctx context.Context, in model.AdminUpdateInput) error
		Delete(ctx context.Context, id uint) error
		GetList(ctx context.Context, in model.AdminGetListInput) (out model.AdminGetListOutput, err error)
		GetUserByUserNamePassword(ctx context.Context, in model.LoginInput) map[string]interface{}
		// i: admin's login input(username, password), o: admin's info (id, name, is_admin, role_ids)
		GetAdminByNamePassword(ctx context.Context, in model.LoginInput) map[string]interface{}
		// i: admin's name, o: admin's info (id, name, is_admin, role_ids)
		GetAdminByNamePasswordRoles(ctx context.Context, in model.LoginInput) map[string]interface{}
		// i: admin's id, o: admin's info (id, name, is_admin, role_ids)
		GetAdminById(ctx context.Context, id uint) (*model.AdminInfo, error)
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
