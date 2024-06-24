// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HomeFriendLink is the golang structure for table home_friend_link.
type HomeFriendLink struct {
	Id         uint64      `json:"id"         ` //
	LinkName   string      `json:"linkName"   ` // 链接名
	LinkUrl    string      `json:"linkUrl"    ` // 链接url
	Sort       uint        `json:"sort"       ` // 排序号
	IsOpen     uint        `json:"isOpen"     ` // 是否开启;0-不开启 1-开启
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
