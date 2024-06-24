// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Literary_Snowlands/internal/dao/internal"
)

// internalUserConsumeLogDao is internal type for wrapping internal DAO implements.
type internalUserConsumeLogDao = *internal.UserConsumeLogDao

// userConsumeLogDao is the data access object for table user_consume_log.
// You can define custom methods on it to extend its functionality as you wish.
type userConsumeLogDao struct {
	internalUserConsumeLogDao
}

var (
	// UserConsumeLog is globally public accessible object for table user_consume_log operations.
	UserConsumeLog = userConsumeLogDao{
		internal.NewUserConsumeLogDao(),
	}
)

// Fill with you ideas below.
