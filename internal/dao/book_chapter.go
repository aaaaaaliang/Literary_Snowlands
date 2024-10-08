// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Literary_Snowlands/internal/dao/internal"
)

// internalBookChapterDao is internal type for wrapping internal DAO implements.
type internalBookChapterDao = *internal.BookChapterDao

// bookChapterDao is the data access object for table book_chapter.
// You can define custom methods on it to extend its functionality as you wish.
type bookChapterDao struct {
	internalBookChapterDao
}

var (
	// BookChapter is globally public accessible object for table book_chapter operations.
	BookChapter = bookChapterDao{
		internal.NewBookChapterDao(),
	}
)

// Fill with you ideas below.
