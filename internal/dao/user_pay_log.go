// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Literary_Snowlands/internal/dao/internal"
)

// internalUserPayLogDao is internal type for wrapping internal DAO implements.
type internalUserPayLogDao = *internal.UserPayLogDao

// userPayLogDao is the data access object for table user_pay_log.
// You can define custom methods on it to extend its functionality as you wish.
type userPayLogDao struct {
	internalUserPayLogDao
}

var (
	// UserPayLog is globally public accessible object for table user_pay_log operations.
	UserPayLog = userPayLogDao{
		internal.NewUserPayLogDao(),
	}
)

// Fill with you ideas below.
