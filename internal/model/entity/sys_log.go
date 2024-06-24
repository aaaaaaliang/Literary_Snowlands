// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure for table sys_log.
type SysLog struct {
	Id         uint64      `json:"id"         ` //
	UserId     uint64      `json:"userId"     ` // 用户id
	Username   string      `json:"username"   ` // 用户名
	Operation  string      `json:"operation"  ` // 用户操作
	Time       uint        `json:"time"       ` // 响应时间
	Method     string      `json:"method"     ` // 请求方法
	Params     string      `json:"params"     ` // 请求参数
	Ip         string      `json:"ip"         ` // IP地址
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
}
