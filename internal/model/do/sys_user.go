// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta     `orm:"table:sys_user, do:true"`
	Id         interface{} //
	Username   interface{} // 用户名
	Password   interface{} // 密码
	Name       interface{} // 真实姓名
	Sex        interface{} // 性别;0-男 1-女
	Birth      *gtime.Time // 出身日期
	Email      interface{} // 邮箱
	Mobile     interface{} // 手机号
	Status     interface{} // 状态;0-禁用 1-正常
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
