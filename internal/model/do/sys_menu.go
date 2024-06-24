// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	Id         interface{} //
	ParentId   interface{} // 父菜单ID;一级菜单为0
	Name       interface{} // 菜单名称
	Url        interface{} // 菜单URL
	Type       interface{} // 类型;0-目录   1-菜单
	Icon       interface{} // 菜单图标
	Sort       interface{} // 排序
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
