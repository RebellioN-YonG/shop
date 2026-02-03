package controller

// todo
// import (
// 	"context"
// "shop/api/frontend"
// 	"shop/internal/model"
// 	"shop/internal/service"

// 	"github.com/gogf/gf/v2/util/gconv"
// )

// User 内容管理
var User = cUser{}

type cUser struct{}

// func (a *cUser) Create(ctx context.Context, req *frontend.RegisterReq) (res *frontend.UserRes, err error) {
// 	input := &model.UserCreateInput{}
// 	input.Name = req.Name
// 	input.Password = req.Password
// 	input.RoleIds = req.RoleIds
// 	input.IsUser = req.IsUser

// 	out, err := service.User().Create(ctx, input)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &frontend.UserRes{UserId: out.UserId}, nil
// }

// // for jwt
// func (c *cUser) Info(ctx context.Context, req *frontend.UserGetInfoReq) (res *frontend.UserGetInfoRes, err error) {
// 	return &frontend.UserGetInfoRes{
// 		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
// 		IdentityKey: service.Auth().IdentityKey,
// 		Payload:     service.Auth().GetPayload(ctx),
// 	}, nil
// }

// func (a *cUser) Delete(ctx context.Context, req *frontend.UserDeleteReq) (res *frontend.UserDeleteRes, err error) {
// 	err = service.User().Delete(ctx, req.UserId)
// 	return
// }

// func (a *cUser) Update(ctx context.Context, req *frontend.UserUpdateReq) (res *frontend.UserUpdateRes, err error) {
// 	err = service.User().Update(ctx, model.UserUpdateInput{
// 		Id: req.UserId,
// 		UserCreateUpdateBase: model.UserCreateUpdateBase{
// 			Name:     req.Name,
// 			Password: req.Password,
// 			RoleIds:  req.RoleIds,
// 			IsUser:   req.IsUser,
// 		},
// 	})
// 	return &frontend.UserUpdateRes{Id: req.UserId}, nil
// }

// func (a *cUser) List(ctx context.Context, req *frontend.UserGetListCommonReq) (res *frontend.UserGetListCommonRes, err error) {
// 	getListRes, err := service.User().GetList(ctx, model.UserGetListInput{
// 		Page: req.Page,
// 		Size: req.Size,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &frontend.UserGetListCommonRes{List: getListRes.List,
// 		Page:  getListRes.Page,
// 		Size:  getListRes.Size,
// 		Total: getListRes.Total}, nil
// }
