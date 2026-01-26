package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminReq struct {
	g.Meta   `path:"/backend/admin/add" tags:"Admin" method:"post" summary:"ur 1st admin api"`
	Name     string `json:"name" v:"required#username cannot be empty"  dc:"name"`
	PassWord string `json:"password" v:"required#password cannot be empty"  dc:"password"`
	RoleIds  string `json:"role_ids"  dc:"role_ids"`
	IsAdmin  int    `json:"is_admin"  dc:"is_admin"`
}

type AdminRes struct {
	AdminId int `json:"admin_id"`
}

type AdminDeleteReq struct {
	g.Meta  `path:"/backend/admin/delete" tags:"Admin" method:"delete" summary:"delete a admin"`
	AdminId uint `json:"admin_id" v:"min:1#select a admin to delete" dc:"admin_id"`
}

type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta  `path:"/backend/admin/update" tags:"Admin" method:"post" summary:"update a admin"`
	AdminId uint   `json:"admin_id" v:"min:1#select a admin to update" dc:"admin_id"`
	PicUrl  string `json:"pic_url" v:"required#cannot be empty" dc:"pic_url"`
	Link    string `json:"link" v:"required#cannot be empty" dc:"link"`
	Sort    int    `json:"sort" dc:"sort"`
}

type AdminUpdateRes struct {
	Id uint `json:"id"`
}

type AdminGetListCommonReq struct {
	g.Meta `path:"/backend/admin/list" tags:"Admin" method:"get" summary:"get admin list"`
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
	g.Meta `path:"/backend/admin/info" tags:"Admin" method:"get" summary:"get admin info"`
}

type AdminGetInfoRes struct {
	AdminId int         `json:"admin_id"`
	IdKey   string      `json:"id_key"`
	PayLoad interface{} `json:"payload"`
}
