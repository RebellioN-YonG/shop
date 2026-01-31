package model

import "github.com/gogf/gf/v2/os/gtime"

type RoleCreateUpdateBase struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// RoleCreateInput 创建角色输入
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput 创建角色输出
type RoleCreateOutput struct {
	RoleId uint `json:"role_id"`
}

// RoleUpdateInput 更新角色输入
type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id uint `json:"id"`
}

// RoleGetListInput 获取角色列表输入
type RoleGetListInput struct {
	Page int `json:"page"`
	Size int `json:"size"`
	Sort int `json:"sort"`
}

// RoleGetListOutput 获取角色列表输出
type RoleGetListOutput struct {
	List  []RoleGetListOutputItem `json:"list" description:"list"`
	Total int                     `json:"total" description:"total"`
	Page  int                     `json:"page" description:"page"`
	Size  int                     `json:"size" description:"size"`
}

// RoleGetListOutputItem 角色列表项
type RoleGetListOutputItem struct {
	Id          uint                 `json:"id"`
	Name        string               `json:"name"`
	Desc        string               `json:"desc"`
	CreateAt    *gtime.Time          `json:"create_time"`
	UpdateAt    *gtime.Time          `json:"update_time"`
	Permissions []RolePermissionItem `json:"permissions"`
}

// RolePermissionItem 角色权限项
type RolePermissionItem struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// RoleAddPermissionInput 添加权限输入
type RoleAddPermissionInput struct {
	RoleId       uint `json:"role_id"`
	PermissionId uint `json:"permission_id"`
}

// RoleAddPermissionOutput 添加权限输出
type RoleAddPermissionOutput struct {
	Id uint `json:"id"`
}

type RoleDeletePermissionInput struct {
	RoleId       uint `json:"role_id"`
	PermissionId uint `json:"permission_id"`
}

type RoleDeletePermissionOutput struct {
}

type RoleSearchOutputItem struct {
	RoleGetListOutputItem
}
