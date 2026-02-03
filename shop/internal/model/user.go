package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserCreateUpdateBase 创建/修改内容基类
type UserCreateUpdateBase struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	UserSalt     string `json:"user_salt"`
	Avatar       string `json:"avatar"`
	Sign         string `json:"sign"`
	SecretAnswer string `json:"secret_answer"`
	Sex          int    `json:"sex"`
	Status       int    `json:"status"`
}

type UserLoginInput struct {
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" DC:"密码"`
}

// RegisterInput 创建内容
type UserRegisterInput struct {
	UserCreateUpdateBase
}

// RegisterOutput 创建内容返回结果
type UserRegisterOutput struct {
	Id uint
}

type UserUpdatePasswordInput struct {
	Password     string `json:"password" v:"required#密码不能为空" dc:"密码"`
	UserSalt     string `json:"user_salt" dc:"盐值"`
	SecretAnswer string `json:"secret_answer" dc:"密保答案"`
}

// UserGetListInput 获取内容列表
type UserGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// UserGetListOutput 查询列表结果
type UserGetListOutput struct {
	List  []UserInfoItem `json:"list" description:"列表"`
	Page  int            `json:"page" description:"分页码"`
	Size  int            `json:"size" description:"分页数量"`
	Total int            `json:"total" description:"数据总数"`
}

// UserInfo
type UserInfoBase struct {
	g.Meta `orm:"table:user_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sign   string `json:"sign"`
	Sex    uint8  `json:"sex"`
	Status uint8  `json:"status"`
}

// UserListItem 主要用于列表展示
type UserInfoItem struct {
	Id        int    `json:"id"` // 自增ID
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Sign      string `json:"sign"`
	Sex       int    `json:"sex"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 修改时间
}
