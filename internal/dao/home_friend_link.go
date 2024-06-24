// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Literary_Snowlands/internal/dao/internal"
)

// internalHomeFriendLinkDao is internal type for wrapping internal DAO implements.
type internalHomeFriendLinkDao = *internal.HomeFriendLinkDao

// homeFriendLinkDao is the data access object for table home_friend_link.
// You can define custom methods on it to extend its functionality as you wish.
type homeFriendLinkDao struct {
	internalHomeFriendLinkDao
}

var (
	// HomeFriendLink is globally public accessible object for table home_friend_link operations.
	HomeFriendLink = homeFriendLinkDao{
		internal.NewHomeFriendLinkDao(),
	}
)

// Fill with you ideas below.
