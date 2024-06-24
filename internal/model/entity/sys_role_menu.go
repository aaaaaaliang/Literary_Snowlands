// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRoleMenu is the golang structure for table sys_role_menu.
type SysRoleMenu struct {
	Id         uint64      `json:"id"         ` //
	RoleId     uint64      `json:"roleId"     ` // 角色ID
	MenuId     uint64      `json:"menuId"     ` // 菜单ID
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
