// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserRole is the golang structure for table sys_user_role.
type SysUserRole struct {
	Id         uint64      `json:"id"         ` //
	UserId     uint64      `json:"userId"     ` // 用户ID
	RoleId     uint64      `json:"roleId"     ` // 角色ID
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
