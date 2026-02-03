package user

import (
	"context"
	"errors"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	utility "shop/utility/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in *model.UserRegisterInput) (out model.UserRegisterOutput, err error) {
	count, err := dao.UserInfo.Ctx(ctx).
		Where(dao.UserInfo.Columns().Name, in.Name).
		Count()
	if err != nil {
		return out, err
	}
	if count > 0 {
		return out, errors.New("User name already exists")
	}
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	in.Status = 1
	//插入数据返回id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).
		Data(in).
		InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserRegisterOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sUser) Delete(ctx context.Context, id uint) error {
	// 删除内容
	_, err := dao.UserInfo.Ctx(ctx).
		Where(dao.UserInfo.Columns().Id, id).
		Delete()
	return err
}

// UpdateStatus 修改状态
func (s *sUser) UpdateStatus(ctx context.Context, id uint, status int) (err error) {
	//更新操作
	_, err = dao.UserInfo.Ctx(ctx).
		Where(dao.UserInfo.Columns().Id, id).
		Data(g.Map{dao.UserInfo.Columns().Status: status}).
		Update()
	return
}

//1.获得*gdb.Model对象，方面后续调用
//2. 实例化响应结构体
//3. 分页查询
//4. 再查询count，判断有无数据
//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
//6. 把查询到的结果赋值到响应结构体中

// GetList 查询内容列表
func (s *sUser) GetList(ctx context.Context, in model.UserGetListInput) (out *model.UserGetListOutput, err error) {
	//1.获得*gdb.Model对象，方便后续调用
	m := dao.UserInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.UserGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size).OrderDesc(dao.UserInfo.Columns().Id)
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		//解决空数据返回[] 而不是返回nil null的问题
		out.List = make([]model.UserInfoItem, 0)
		return out, err
	}

	// 查询管理员列表
	var list []entity.UserInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}

	// 构建输出列表
	out.List = make([]model.UserInfoItem, 0, len(list))
	for _, user := range list {
		out.List = append(out.List, model.UserInfoItem{
			Id:        user.Id,
			Name:      user.Name,
			Avatar:    user.Avatar,
			Sign:      user.Sign,
			Sex:       user.Sex,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}
	return
}

// GetById 根据ID获取管理员信息
func (s *sUser) GetById(ctx context.Context, id int) (*model.UserInfoItem, error) {
	var userInfo entity.UserInfo
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Id, id).Scan(&userInfo)
	if err != nil {
		return nil, err
	}
	if userInfo.Id == 0 {
		return nil, nil
	}
	return &model.UserInfoItem{
		Id:        userInfo.Id,
		Name:      userInfo.Name,
		Avatar:    userInfo.Avatar,
		Sign:      userInfo.Sign,
		Sex:       userInfo.Sex,
		Status:    userInfo.Status,
		CreatedAt: userInfo.CreatedAt.String(),
		UpdatedAt: userInfo.UpdatedAt.String(),
	}, nil
}
