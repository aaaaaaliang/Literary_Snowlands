// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorCode is the golang structure for table author_code.
type AuthorCode struct {
	Id           uint64      `json:"id"           ` // 主键
	InviteCode   string      `json:"inviteCode"   ` // 邀请码
	ValidityTime *gtime.Time `json:"validityTime" ` // 有效时间
	IsUsed       uint        `json:"isUsed"       ` // 是否使用过;0-未使用 1-使用过
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 更新时间
}
