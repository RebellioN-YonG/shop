package model

import "github.com/gogf/gf/v2/os/gtime"

type PermissionCreateUpdateBase struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// PermissionCreateInput 创建权限输入
type PermissionCreateInput struct {
	PermissionCreateUpdateBase
}

// PermissionCreateOutput 创建权限输出
type PermissionCreateOutput struct {
	PermissionId uint `json:"permission_id"`
}

// PermissionUpdateInput 更新权限输入
type PermissionUpdateInput struct {
	PermissionCreateUpdateBase
	Id uint `json:"id"`
}

// PermissionGetListInput 获取权限列表输入
type PermissionGetListInput struct {
	Page int `json:"page"`
	Size int `json:"size"`
	Sort int `json:"sort"`
}

// PermissionGetListOutput 获取权限列表输出
type PermissionGetListOutput struct {
	List  []PermissionGetListOutputItem `json:"list" pathription:"list"`
	Total int                           `json:"total" pathription:"total"`
	Page  int                           `json:"page" pathription:"page"`
	Size  int                           `json:"size" pathription:"size"`
}

// PermissionGetListOutputItem 权限列表项
type PermissionGetListOutputItem struct {
	Id       uint        `json:"id"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	CreateAt *gtime.Time `json:"create_time"`
	UpdateAt *gtime.Time `json:"update_time"`
}

type PermissionSearchOutputItem struct {
	PermissionGetListOutputItem
}
