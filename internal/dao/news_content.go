// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Literary_Snowlands/internal/dao/internal"
)

// internalNewsContentDao is internal type for wrapping internal DAO implements.
type internalNewsContentDao = *internal.NewsContentDao

// newsContentDao is the data access object for table news_content.
// You can define custom methods on it to extend its functionality as you wish.
type newsContentDao struct {
	internalNewsContentDao
}

var (
	// NewsContent is globally public accessible object for table news_content operations.
	NewsContent = newsContentDao{
		internal.NewNewsContentDao(),
	}
)

// Fill with you ideas below.
