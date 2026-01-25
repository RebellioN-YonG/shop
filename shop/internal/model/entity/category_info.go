// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure for table category_info.
type CategoryInfo struct {
	Id          int         `json:"id"          orm:"id"          description:""`
	ParentId    int         `json:"parentId"    orm:"parent_id"   description:"父级id"`
	Name        string      `json:"name"        orm:"name"        description:"分类名称"`
	PicUrl      string      `json:"picUrl"      orm:"pic_url"     description:"分类图标"`
	Level       int         `json:"level"       orm:"level"       description:"等级 默认1级分类"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序值"`
	Status      int         `json:"status"      orm:"status"      description:"状态 1:启用 0:禁用"`
	Description string      `json:"description" orm:"description" description:"分类描述"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:""`
}
