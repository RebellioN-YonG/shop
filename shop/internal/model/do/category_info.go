// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure of table category_info for DAO operations like Where/Data.
type CategoryInfo struct {
	g.Meta      `orm:"table:category_info, do:true"`
	Id          any         //
	ParentId    any         // 父级id
	Name        any         // 分类名称
	PicUrl      any         // 分类图标
	Level       any         // 等级 默认1级分类
	Sort        any         // 排序值
	Status      any         // 状态 1:启用 0:禁用
	Description any         // 分类描述
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}
