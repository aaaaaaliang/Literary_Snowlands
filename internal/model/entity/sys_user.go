// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id         uint64      `json:"id"         ` //
	Username   string      `json:"username"   ` // 用户名
	Password   string      `json:"password"   ` // 密码
	Name       string      `json:"name"       ` // 真实姓名
	Sex        uint        `json:"sex"        ` // 性别;0-男 1-女
	Birth      *gtime.Time `json:"birth"      ` // 出身日期
	Email      string      `json:"email"      ` // 邮箱
	Mobile     string      `json:"mobile"     ` // 手机号
	Status     uint        `json:"status"     ` // 状态;0-禁用 1-正常
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
