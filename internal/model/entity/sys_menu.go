// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id         uint64      `json:"id"         ` //
	ParentId   uint64      `json:"parentId"   ` // 父菜单ID;一级菜单为0
	Name       string      `json:"name"       ` // 菜单名称
	Url        string      `json:"url"        ` // 菜单URL
	Type       uint        `json:"type"       ` // 类型;0-目录   1-菜单
	Icon       string      `json:"icon"       ` // 菜单图标
	Sort       uint        `json:"sort"       ` // 排序
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
