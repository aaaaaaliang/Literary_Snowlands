// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Literary_Snowlands/internal/dao/internal"
)

// internalUserCommentReplyDao is internal type for wrapping internal DAO implements.
type internalUserCommentReplyDao = *internal.UserCommentReplyDao

// userCommentReplyDao is the data access object for table user_comment_reply.
// You can define custom methods on it to extend its functionality as you wish.
type userCommentReplyDao struct {
	internalUserCommentReplyDao
}

var (
	// UserCommentReply is globally public accessible object for table user_comment_reply operations.
	UserCommentReply = userCommentReplyDao{
		internal.NewUserCommentReplyDao(),
	}
)

// Fill with you ideas below.