package model

import "github.com/gogf/gf/v2/os/gtime"

type AdminCreateUpdateBase struct {
	Name     string `json:"name" v:"required#username cannot be empty"  dc:"name"`
	PassWord string `json:"password" v:"required#password cannot be empty"  dc:"password"`
	RoleIds  string `json:"role_ids"  dc:"role_ids"`
	UserSalt string `json:"user_salt"  dc:"user_salt"`
	IsAdmin  int    `json:"is_admin"  dc:"is_admin"`
}

// create admin input
type AdminCreateInput struct {
	AdminCreateUpdateBase
}

// create admin output
type AdminCreateOutput struct {
	AdminId int `json:"admin_id"`
}

type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}

type AdminGetListInput struct {
	Page int
	Size int
	Sort int
}

type AdminGetListOutput struct {
	List  []AdminGetListOutputItem `json:"list" description:"list"`
	Total int                      `json:"total" description:"total"`
	Page  int                      `json:"page" description:"page"`
	Size  int                      `json:"size" description:"size"`
}

type AdminGetListOutputItem struct {
	Id         uint        `json:"id"`
	Name       string      `json:"name"`
	PassWord   string      `json:"password"`
	RoleIds    string      `json:"role_ids"`
	RoleIdList []int       `json:"role_id_list"`
	IsAdmin    int         `json:"is_admin"`
	CreateAt   *gtime.Time `json:"create_time"`
	UpdateAt   *gtime.Time `json:"update_time"`
}

type AdminSearchOutputItem struct {
	AdminGetListOutputItem
}

type AdminInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	RoleIds string `json:"role_ids"`
	IsAdmin int    `json:"is_admin"`
}

// type AdminSearchInput struct {
// 	Key        string // search key
// 	Type       string // search type
// 	CategoryId uint   // category id
// 	Page       int    // page
// 	Size       int    // size: max 50
// 	Sort       int    // sort: 0 newest, 1 activity, 2 hot
// }

// type AdminSearchOutput struct {
// 	List  []AdminGetListOutput `json:"list"`
// 	Stats map[string]int       `json:"stats"`
// 	Page  int                  `json:"page"`
// 	Size  int                  `json:"size"`
// 	Total int                  `json:"total"`
// }
