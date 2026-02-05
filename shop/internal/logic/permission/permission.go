package permission

import (
	"context"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

// Create 创建角色
func (s *sPermission) Create(ctx context.Context, in *model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	// 插入并返回 ID
	lastInsertId, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: uint(lastInsertId)}, nil
}

// Update 更新角色
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	_, err := dao.PermissionInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.PermissionInfo.Columns().Id).
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	return err
}

// Delete 删除角色
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	// 删除角色
	_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
		dao.PermissionInfo.Columns().Id: id,
	}).Unscoped().Delete()
	return err
}

// GetList 获取角色列表（含权限）
func (s *sPermission) GetList(ctx context.Context, in *model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	// 1. 获取 gdb.Model
	m := dao.PermissionInfo.Ctx(ctx)
	// 2. 初始化返回结构
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 3. 分页查询
	listModel := m.Page(in.Page, in.Size)

	// 4. 查询角色列表
	var list []entity.PermissionInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	if len(list) == 0 {
		out.List = make([]model.PermissionGetListOutputItem, 0)
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// GetPathsByRoleIds 根据角色ID获取权限列表
func (s *sPermission) GetPathsByRoleIds(ctx context.Context, RoleIds []int) ([]string, error) {
	if len(RoleIds) == 0 {
		return []string{}, nil
	}
	// 根据 RoleId 查询 RolePermissionInfo 关联表获取 PermissionInfo
	var RolePermissions []entity.RolePermissionInfo
	err := dao.RolePermissionInfo.Ctx(ctx).
		WhereIn(dao.RolePermissionInfo.Columns().RoleId, RoleIds).
		Scan(&RolePermissions)
	if err != nil {
		return nil, err
	}

	if len(RolePermissions) == 0 {
		g.Log().Debug(ctx, "GetPathsByRoleIds - no role permissions found for RoleIds:", RoleIds)
		return []string{}, nil
	}

	// 提取 PermissionId 后去重
	permissionIdMap := make(map[int]bool)
	for _, rp := range RolePermissions {
		permissionIdMap[rp.PermissionId] = true
	}

	// 转换回 slice
	permissionIds := make([]int, 0, len(permissionIdMap))
	for id := range permissionIdMap {
		permissionIds = append(permissionIds, id)
	}
	g.Log().Debug(ctx, "GetPathsByRoleIds - permissionIds:", permissionIds)

	// 根据 permissionIds 查询权限详情以获取 path
	var permissionList []entity.PermissionInfo
	err = dao.PermissionInfo.Ctx(ctx).
		WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).
		Scan(&permissionList)
	if err != nil {
		return nil, err
	}

	// 提取所有 path 构建返回结果
	paths := make([]string, 0, len(permissionList))
	for _, p := range permissionList {
		paths = append(paths, p.Path)
	}
	g.Log().Debug(ctx, "GetPathsByRoleIds - paths:", paths)
	return paths, nil
}
