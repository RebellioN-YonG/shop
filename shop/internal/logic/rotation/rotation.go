package logic

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// html is not allowed
	if err = ghtml.SpecialCharsMapOrStruct(&in); err != nil {
		return out, err
	}
	lastInsertId, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertId)}, err
}

func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// html is not allowed
		if err := ghtml.SpecialCharsMapOrStruct(&in); err != nil {
			return err
		}
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}
