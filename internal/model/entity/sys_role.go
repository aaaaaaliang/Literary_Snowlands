// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id         uint64      `json:"id"         ` //
	RoleName   string      `json:"roleName"   ` // 角色名称
	RoleSign   string      `json:"roleSign"   ` // 角色标识
	Remark     string      `json:"remark"     ` // 备注
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
