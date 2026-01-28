package admin

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	utility "shop/utility/utils"
	"strconv"
	"strings"

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

func (s *sAdmin) Create(ctx context.Context, in *model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// html is not allowed
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	// process salt and password
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	// return id after insert
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
		}).Delete()
		return err
	})
}

func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out model.AdminGetListOutput, err error) {
	// 1. get gdb.Model
	m := dao.AdminInfo.Ctx(ctx)
	// 2.instantiate resp
	out = model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 3.  query by pages and size
	litsModel := m.Page(in.Page, in.Size)
	// 4. get total count
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 5. query admin list
	var list []entity.AdminInfo
	if err := litsModel.Scan(&list); err != nil {
		return out, err
	}

	// 6. make output list and convert role_ids to a array
	out.List = make([]model.AdminGetListOutputItem, 0, len(list))
	for _, admin := range list {
		// convert role_ids to a array
		var roleIdArray []int
		if admin.RoleIds != "" {
			roleIdStrs := strings.Split(admin.RoleIds, ",")
			for _, idStr := range roleIdStrs {
				idStr = strings.TrimSpace(idStr)
				if idStr != "" {
					if id, err := strconv.Atoi(idStr); err == nil {
						roleIdArray = append(roleIdArray, id)
					}
				}
			}
		}
		out.List = append(out.List, model.AdminGetListOutputItem{
			Id:         uint(admin.Id),
			Name:       admin.Name,
			RoleIds:    admin.RoleIds,
			RoleIdList: roleIdArray,
			IsAdmin:    admin.IsAdmin,
			CreateAt:   admin.CreatedAt,
			UpdateAt:   admin.UpdatedAt,
		})
	}
	return
}

// i: user's login input(username, password), o: user's info (id, name)
func (s *sAdmin) GetUserByUserNamePassword(ctx context.Context, in model.LoginInput) map[string]interface{} {
	// 1. check if password is correct
	var adminInfo entity.AdminInfo
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&adminInfo)
	if err != nil || adminInfo.Id == 0 {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	}
	return g.Map{
		"id":   adminInfo.Id,
		"name": adminInfo.Name,
	}

}

// i: admin's login input(username, password), o: admin's info (id, name, is_admin, role_ids)
func (s *sAdmin) GetAdminByNamePassword(ctx context.Context, in model.LoginInput) map[string]interface{} {
	// 1. check if password is correct
	var adminInfo entity.AdminInfo
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&adminInfo)
	if err != nil || adminInfo.Id == 0 {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	}
	return g.Map{
		"id":       adminInfo.Id,
		"name":     adminInfo.Name,
		"is_admin": adminInfo.IsAdmin,
		"role_ids": adminInfo.RoleIds,
	}
}

// i: admin's name, o: admin's info (id, name, is_admin, role_ids)
func (s *sAdmin) GetAdminByNamePasswordRoles(ctx context.Context, in model.LoginInput) map[string]interface{} {
	var adminInfo entity.AdminInfo
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	}
	return g.Map{
		"id":       adminInfo.Id,
		"name":     adminInfo.Name,
		"is_admin": adminInfo.IsAdmin,
		"role_ids": adminInfo.RoleIds,
	}
}

// i: admin's id, o: admin's info (id, name, is_admin, role_ids)
func (s *sAdmin) GetAdminById(ctx context.Context, id uint) (*model.AdminInfo, error) {
	var adminInfo entity.AdminInfo
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Id, id).Scan(&adminInfo)
	if err != nil || adminInfo.Id == 0 {
		return nil, err
	}
	if adminInfo.Id == 0 {
		return nil, nil
	}
	return &model.AdminInfo{
		Id:      adminInfo.Id,
		Name:    adminInfo.Name,
		IsAdmin: adminInfo.IsAdmin,
		RoleIds: adminInfo.RoleIds,
	}, nil
}
