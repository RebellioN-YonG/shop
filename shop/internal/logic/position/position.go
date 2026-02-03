package position

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error) {
	// html is not allowed
	if err = ghtml.SpecialCharsMapOrStruct(&in); err != nil {
		return out, err
	}
	lastInsertId, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PositionCreateOutput{PositionId: int(lastInsertId)}, err
}

func (s *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// html is not allowed
		if err := ghtml.SpecialCharsMapOrStruct(&in); err != nil {
			return err
		}
		_, err := dao.PositionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PositionInfo.Columns().Id).
			Where(dao.PositionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

func (s *sPosition) Delete(ctx context.Context, id uint) error {
	return dao.PositionInfo.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// del
		_, err := dao.PositionInfo.Ctx(ctx).Where(g.Map{
			dao.PositionInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

func (s *sPosition) GetList(ctx context.Context, in model.PositionGetListInput) (out model.PositionGetListOutput, err error) {
	// 1. get gdb.Model
	m := dao.PositionInfo.Ctx(ctx)
	//
	out = model.PositionGetListOutput{
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
	out.List = make([]model.PositionGetListOutputItem, 0, in.Size)
	if err := litsModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
