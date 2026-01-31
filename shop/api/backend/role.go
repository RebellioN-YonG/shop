package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RoleReq struct {
	g.Meta `path:"/role/add" tags:"Role" method:"post" desc:"add role"`
	Name   string `json:"name" v:"required#rolename cannot be empty"  dc:"role name"`
	Desc   string `json:"desc" dc:"role description"`
}

type RoleRes struct {
	RoleId uint `json:"role_id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" tags:"Role" method:"delete" summary:"delete a role"`
	Id     uint `json:"id" v:"required#role id cannot be empty" dc:"role id"`
}

type RoleDeleteRes struct{}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" tags:"Role" method:"post" summary:"update a role"`
	Id     uint   `json:"id" v:"required#role id cannot be empty" dc:"role id"`
	Name   string `json:"name" v:"required#rolename cannot be empty"  dc:"role name"`
	Desc   string `json:"desc" dc:"role description"`
}

type RoleUpdateRes struct {
	Id uint `json:"id"`
}

type RoleGetListCommonReq struct {
	g.Meta `path:"/role/list" tags:"Role" method:"get" summary:"get role list"`
	CommonPaginationReq
}

type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"role list"`
	Page  int         `json:"page" description:"page number"`
	Size  int         `json:"size" description:"size of page"`
	Total int         `json:"total" description:"total amount"`
}

type AddPermissionReq struct {
	g.Meta       `path:"/role/add/permission" tags:"Role" method:"post" summary:"add permission to role"`
	RoleId       uint `json:"role_id" v:"required#role id cannot be empty" dc:"role id"`
	PermissionId uint `json:"permission_id" v:"required#permission id cannot be empty" dc:"permission id"`
}

type AddPermissionRes struct {
	Id uint `json:"id"`
}

type AddPermissionsReq struct {
	g.Meta        `path:"/role/add/permissions" tags:"Role" method:"post" summary:"add permission to role batch"`
	RoleId        uint   `json:"role_id" v:"required#role id cannot be empty" dc:"role id"`
	PermissionIds []uint `json:"permission_ids" dc:"permission id"`
}

type AddPermissionsRes struct{}

type DeletePermissionReq struct {
	g.Meta       `path:"/role/delete/permission" tags:"Role" method:"delete" summary:"delete permission from role"`
	RoleId       uint `json:"role_id" v:"required#role id cannot be empty" dc:"role id"`
	PermissionId uint `json:"permission_id" v:"required#permission id cannot be empty" dc:"permission id"`
}
type DeletePermissionsReq struct {
	g.Meta        `path:"/role/delete/permissions" tags:"Role" method:"delete" summary:"delete permission from role batch"`
	RoleId        uint   `json:"role_id" v:"required#role id cannot be empty" dc:"role id"`
	PermissionIds []uint `json:"permission_ids" dc:"permission id"`
}

type DeletePermissionRes struct{}
type DeletePermissionsRes struct{}
