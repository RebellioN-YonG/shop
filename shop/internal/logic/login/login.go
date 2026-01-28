package login

import (
	"context"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	utility "shop/utility/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// for jwt
func (s *sLogin) Login(ctx context.Context, in model.LoginInput) error {
	var adminInfo entity.AdminInfo
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	gutil.Dump("encrypted password: ", utility.EncryptPassword(in.Name, adminInfo.UserSalt))
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("username or password error")
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// auto update Login
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}

func (s *sLogin) Logout(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}
