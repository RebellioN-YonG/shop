package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PermissionCreateUpdateBase struct {
	Name string `json:"name" v:"required#name cannot be empty" dc:"permission name"`
	Path string `json:"path"`
}

type PermissionReq struct {
	g.Meta `path:"/permission/add" tags:"Permission" method:"post" desc:"add permission"`
	PermissionCreateUpdateBase
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/permission/delete" tags:"Permission" method:"delete" summary:"delete a permission"`
	Id     uint `json:"id" v:"required#permission id cannot be empty" dc:"permission id"`
}

type PermissionDeleteRes struct{}

type PermissionUpdateReq struct {
	g.Meta `path:"/permission/update" tags:"Permission" method:"post" summary:"update a permission"`
	Id     uint `json:"id" v:"required#id cannot be empty" dc:"id"`
	PermissionCreateUpdateBase
}

type PermissionUpdateRes struct {
	Id uint `json:"id"`
}

type PermissionGetListCommonReq struct {
	g.Meta `path:"/permission/list" tags:"Permission" method:"get" summary:"get permission list"`
	CommonPaginationReq
}

type PermissionGetListCommonRes struct {
	List  interface{} `json:"list" description:"permission list"`
	Page  int         `json:"page" description:"page number"`
	Size  int         `json:"size" description:"size of page"`
	Total int         `json:"total" description:"total amount"`
}
