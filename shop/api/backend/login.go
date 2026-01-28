package backend

import (
	"shop/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginDoReq struct {
	g.Meta   `path:"/login" tags:"Login" method:"post" summary:"execute login request"`
	Name     string `json:"name" v:"required#cannot be empty" dc:"name"`
	Password string `json:"password" v:"required#cannot be empty" dc:"password"`
}

// for jwt auth
type LoginDoRes struct {
	// Info   interface{} `json:"info"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// for gtoken
type LoginRes struct {
	Type       string                  `json:"type"`
	Token      string                  `json:"token"`
	ExpireIn   int                     `json:"expire_in"`
	IsAdmin    int                     `json:"is_admin"`   // is super admin
	RoleIds    string                  `json:"role_ids"`   // role ids
	Permission []entity.PermissionInfo `json:"permission"` // permission info list
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post"`
}

type LogoutRes struct {
}
