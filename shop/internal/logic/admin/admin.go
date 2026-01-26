package logic

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
	util "shop/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// html is not allowed
	if err = ghtml.SpecialCharsMapOrStruct(&in); err != nil {
		return out, err
	}
	// process salt and password
	UserSalt := grand.S(10)
	in.PassWord = util.EncryptPassword(in.PassWord, UserSalt)
	in.UserSalt = UserSalt
	lastInsertId, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: int(lastInsertId)}, err
}

func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// html is not allowed
		if err := ghtml.SpecialCharsMapOrStruct(&in); err != nil {
			return err
		}
		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id).
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	return dao.AdminInfo.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// del
		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out model.AdminGetListOutput, err error) {
	// 1. get gdb.Model
	m := dao.AdminInfo.Ctx(ctx)
	//
	out = model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// query by pages and size
	litsModel := m.Page(in.Page, in.Size)
	//
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	out.List = make([]model.AdminGetListOutputItem, 0, in.Size)
	if err := litsModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
