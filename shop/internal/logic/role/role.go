package role

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Create 创建角色
func (s *sRole) Create(ctx context.Context, in *model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	// 插入并返回 ID
	lastInsertId, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: uint(lastInsertId)}, nil
}

// Update 更新角色
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	_, err := dao.RoleInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.RoleInfo.Columns().Id).
		Where(dao.RoleInfo.Columns().Id, in.Id).
		Update()
	return err
}

// Delete 删除角色
func (s *sRole) Delete(ctx context.Context, id uint) error {
	// 删除角色
	_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
		dao.RoleInfo.Columns().Id: id,
	}).Unscoped().Delete()
	return err
}

// GetList 获取角色列表（含权限）
func (s *sRole) GetList(ctx context.Context, in *model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	// 1. 获取 gdb.Model
	m := dao.RoleInfo.Ctx(ctx)
	// 2. 初始化返回结构
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 3. 分页查询
	listModel := m.Page(in.Page, in.Size)

	// 4. 查询角色列表
	var list []entity.RoleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	if len(list) == 0 {
		out.List = make([]model.RoleGetListOutputItem, 0)
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// 6. 构建返回列表，包含每个角色的权限
	out.List = make([]model.RoleGetListOutputItem, 0, len(list))
	for _, role := range list {
		// 查询该角色的权限列表
		permissions, _ := s.GetPermissionsByRoleId(ctx, uint(role.Id))

		out.List = append(out.List, model.RoleGetListOutputItem{
			Id:          uint(role.Id),
			Name:        role.Name,
			Desc:        role.Desc,
			CreateAt:    role.CreatedAt,
			UpdateAt:    role.UpdatedAt,
			Permissions: permissions,
		})
	}
	return
}

// GetPermissionsByRoleId 根据角色ID获取权限列表
func (s *sRole) GetPermissionsByRoleId(ctx context.Context, roleId uint) ([]model.RolePermissionItem, error) {
	var permissions []model.RolePermissionItem

	// 查询角色权限关联表
	var rolePermissions []entity.RolePermissionInfo
	err := dao.RolePermissionInfo.Ctx(ctx).
		Where(dao.RolePermissionInfo.Columns().RoleId, roleId).
		Scan(&rolePermissions)
	if err != nil {
		return permissions, err
	}

	if len(rolePermissions) == 0 {
		return permissions, nil
	}

	// 获取权限ID列表
	permissionIds := make([]int, 0, len(rolePermissions))
	for _, rp := range rolePermissions {
		permissionIds = append(permissionIds, rp.PermissionId)
	}

	// 查询权限详情
	var permissionList []entity.PermissionInfo
	err = dao.PermissionInfo.Ctx(ctx).
		WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).
		Scan(&permissionList)
	if err != nil {
		return permissions, err
	}

	// 构建返回结果
	for _, p := range permissionList {
		permissions = append(permissions, model.RolePermissionItem{
			Id:   uint(p.Id),
			Name: p.Name,
			Path: p.Path,
		})
	}
	return permissions, nil
}

// AddPermission 为角色添加单个权限
func (s *sRole) AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error) {
	// 检查是否已存在
	count, err := dao.RolePermissionInfo.Ctx(ctx).
		Where(g.Map{
			dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
			dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
		}).Count()
	if err != nil {
		return out, err
	}
	if count > 0 {
		// 已存在，返回成功但不重复添加
		return out, nil
	}

	// 插入新记录
	lastInsertId, err := dao.RolePermissionInfo.Ctx(ctx).Data(g.Map{
		dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
		dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
	}).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleAddPermissionOutput{Id: uint(lastInsertId)}, nil
}

// AddPermissions 批量为角色添加权限
func (s *sRole) AddPermissions(ctx context.Context, roleId uint, permissionIds []uint) error {
	return dao.RolePermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, permissionId := range permissionIds {
			// 检查是否已存在
			count, err := dao.RolePermissionInfo.Ctx(ctx).
				Where(g.Map{
					dao.RolePermissionInfo.Columns().RoleId:       roleId,
					dao.RolePermissionInfo.Columns().PermissionId: permissionId,
				}).Count()
			if err != nil {
				return err
			}
			if count > 0 {
				continue // 已存在则跳过
			}

			// 插入新记录
			_, err = dao.RolePermissionInfo.Ctx(ctx).Data(g.Map{
				dao.RolePermissionInfo.Columns().RoleId:       roleId,
				dao.RolePermissionInfo.Columns().PermissionId: permissionId,
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// DeletePermission 删除角色权限
func (s *sRole) DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error {
	_, err := dao.RolePermissionInfo.Ctx(ctx).Where(
		g.Map{
			dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
			dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
		}).Unscoped().Delete()
	return err
}
func (s *sRole) DeletePermissions(ctx context.Context, roleId uint, permissionIds []uint) error {
	_, err := dao.RolePermissionInfo.Ctx(ctx).
		Where(dao.RolePermissionInfo.Columns().RoleId, roleId).
		WhereIn(dao.RolePermissionInfo.Columns().PermissionId, permissionIds).
		Unscoped().Delete()
	return err
}
