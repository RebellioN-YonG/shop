package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserPromoteToAdminReq struct {
	g.Meta   `path:"/user/add" tags:"User" method:"post" summary:"You first user api"`
	Id       string `json:"id" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色ids"`
}

type UserPromoteToAdminRes struct {
	AdminId int `json:"admin_id"`
}

type UserDeleteReq struct {
	g.Meta `path:"/user/delete" method:"delete" tags:"管理员" summary:"删除管理员接口"`
	Id     uint `json:"id" v:"min:1#请选择需要删除的管理员" dc:"管理员id"`
}

type UserDeleteRes struct {
}

type UserUpdateStatusReq struct {
	g.Meta `path:"/user/update" method:"post" tags:"管理员" summary:"修改管理员接口"`
	Id     uint `json:"id" v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Status int  `json:"status" v:"required|in:0,1" dc:"状态, 0禁用 1启用"`
}

type UserUpdateStatusRes struct {
}

type UserGetListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"管理员" summary:"管理员列表接口"`
	CommonPaginationReq
}
type UserGetListRes struct {
	List  []UserListItem `json:"list" description:"列表"`
	Page  int            `json:"page" description:"分页码"`
	Size  int            `json:"size" description:"分页数量"`
	Total int            `json:"total" description:"数据总数"`
}

type UserListItem struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Avatar   string      `json:"avatar"`
	Sex      int         `json:"sex"`
	Status   int         `json:"status"`
	Sign     string      `json:"sign"`
	CreateAt *gtime.Time `json:"create_at"`
	UpdateAt *gtime.Time `json:"update_at"`
}
