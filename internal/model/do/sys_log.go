// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure of table sys_log for DAO operations like Where/Data.
type SysLog struct {
	g.Meta     `orm:"table:sys_log, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户id
	Username   interface{} // 用户名
	Operation  interface{} // 用户操作
	Time       interface{} // 响应时间
	Method     interface{} // 请求方法
	Params     interface{} // 请求参数
	Ip         interface{} // IP地址
	CreateTime *gtime.Time // 创建时间
}
