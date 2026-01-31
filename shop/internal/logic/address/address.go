package address

import (
	"context"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
)

type sAddress struct{}

func init() {
	service.RegisterAddress(&sAddress{})
}

// Add 添加地址
func (s *sAddress) Add(ctx context.Context, in *model.AddAddressInput) (out *model.AddAddressOutput, err error) {
	id, err := dao.AddressInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.AddAddressOutput{Id: int(id)}, nil
}

// Update 更新地址
func (s *sAddress) Update(ctx context.Context, in *model.UpdateAddressInput) (err error) {
	if _, err = dao.AddressInfo.Ctx(ctx).
		Data(in).
		OmitEmpty().
		Where(dao.AddressInfo.Columns().Id, in.Id).
		Update(); err != nil {
		return err
	}
	return nil
}

// Delete 删除地址
func (s *sAddress) Delete(ctx context.Context, id int) (err error) {
	if _, err = dao.AddressInfo.Ctx(ctx).
		Where(dao.AddressInfo.Columns().Id, id).
		Delete(); err != nil {
		return err
	}
	return nil
}

// Page 分页查询地址
func (s *sAddress) Page(ctx context.Context, in *model.PageAddressInput) (out *model.PageAddressOutput, err error) {
	out = &model.PageAddressOutput{}
	out.Page = in.Page
	out.Size = in.Size

	m := dao.AddressInfo.Ctx(ctx)

	// 获取总数
	out.Total, err = m.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	var list []model.CityAddressListOutputItem
	err = m.Page(in.Page, in.Size).Scan(&list)
	if err != nil {
		return nil, err
	}

	out.List = list
	return out, nil
}

// GetCityList 获取城市地址列表（不分页）
func (s *sAddress) GetCityList(ctx context.Context) (out *model.CityAddressListOutput, err error) {
	out = &model.CityAddressListOutput{}

	err = dao.AddressInfo.Ctx(ctx).
		Where(dao.AddressInfo.Columns().Pid, consts.ProvincePid).
		WithAll().Scan(&out.List)
	if err != nil {
		return nil, err
	}
	return out, nil
}
