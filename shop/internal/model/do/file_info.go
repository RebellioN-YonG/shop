// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure of table file_info for DAO operations like Where/Data.
type FileInfo struct {
	g.Meta    `orm:"table:file_info, do:true"`
	Id        any         //
	Name      any         // 图片名称
	Src       any         // 本地文件存储路径
	Url       any         // URL地址
	UserId    any         // 用户id
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
