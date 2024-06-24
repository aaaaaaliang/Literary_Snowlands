// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorInfo is the golang structure for table author_info.
type AuthorInfo struct {
	Id            uint64      `json:"id"            ` // 主键
	UserId        uint64      `json:"userId"        ` // 用户ID
	InviteCode    string      `json:"inviteCode"    ` // 邀请码
	PenName       string      `json:"penName"       ` // 笔名
	TelPhone      string      `json:"telPhone"      ` // 手机号码
	ChatAccount   string      `json:"chatAccount"   ` // QQ或微信账号
	Email         string      `json:"email"         ` // 电子邮箱
	WorkDirection uint        `json:"workDirection" ` // 作品方向;0-男频 1-女频
	Status        uint        `json:"status"        ` // 0：正常;1-封禁
	CreateTime    *gtime.Time `json:"createTime"    ` // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    ` // 更新时间
}
