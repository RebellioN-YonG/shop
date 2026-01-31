package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminReq struct {
	g.Meta   `path:"/admin/add" tags:"Admin" method:"post" summary:"ur 1st admin api"`
	Name     string `json:"name" v:"required#username cannot be empty"  dc:"name"`
	Password string `json:"password" v:"required#password cannot be empty"  dc:"password"`
	RoleIds  string `json:"role_ids"  dc:"role_ids"`
	IsAdmin  int    `json:"is_admin"  dc:"is_admin"`
}

type AdminRes struct {
	AdminId int `json:"admin_id"`
}

type AdminDeleteReq struct {
	g.Meta  `path:"/admin/delete" tags:"Admin" method:"delete" summary:"delete a admin"`
	AdminId uint `json:"admin_id" v:"min:1#select a admin to delete" dc:"admin's id"`
}

type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/admin/update" tags:"Admin" method:"post" summary:"update a admin"`
	AdminId  uint   `json:"admin_id" v:"min:1#select a admin to update" dc:"admin's id"`
	Name     string `json:"name" v:"required#username cannot be empty"  dc:"name"`
	Password string `json:"password" v:"required#password cannot be empty"  dc:"password"`
	RoleIds  string `json:"role_ids"  dc:"role_ids"`
	IsAdmin  int    `json:"is_admin"  dc:"is_admin"`
}

type AdminUpdateRes struct {
	Id uint `json:"id"`
}

type AdminGetListCommonReq struct {
	g.Meta `path:"/admin/list" tags:"Admin" method:"get" summary:"get admin list"`
	Sort   int `json:"sort" in:"query" dc:"sort"`
	CommonPaginationReq
}

type AdminGetListCommonRes struct {
	List  interface{} `json:"list" description:"admin list"`
	Page  int         `json:"page" description:"page number"`
	Size  int         `json:"size" description:"size of page"`
	Total int         `json:"total" description:"total amount"`
}

type AdminGetInfoReq struct {
	g.Meta `path:"/admin/info" tags:"Admin" method:"get" summary:"get admin info"`
}

type AdminGetInfoRes struct {
	Id      int         `json:"id"`
	IdKey   string      `json:"id_key"`
	PayLoad interface{} `json:"payload"`
}
