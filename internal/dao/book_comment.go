// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Literary_Snowlands/internal/dao/internal"
)

// internalBookCommentDao is internal type for wrapping internal DAO implements.
type internalBookCommentDao = *internal.BookCommentDao

// bookCommentDao is the data access object for table book_comment.
// You can define custom methods on it to extend its functionality as you wish.
type bookCommentDao struct {
	internalBookCommentDao
}

var (
	// BookComment is globally public accessible object for table book_comment operations.
	BookComment = bookCommentDao{
		internal.NewBookCommentDao(),
	}
)

// Fill with you ideas below.
