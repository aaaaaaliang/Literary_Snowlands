// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserReadHistory is the golang structure for table user_read_history.
type UserReadHistory struct {
	Id           uint64      `json:"id"           ` // 主键
	UserId       uint64      `json:"userId"       ` // 用户ID
	BookId       uint64      `json:"bookId"       ` // 小说ID
	PreContentId uint64      `json:"preContentId" ` // 上一次阅读的章节内容表ID
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间;
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 更新时间;
}
