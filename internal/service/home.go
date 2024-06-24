// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/entity"
	"context"
)

type (
	IHome interface {
		FindRecommendBook(ctx context.Context) (recommendBooksInfo []model.RecommendedBooks, err error)
		QueryFriendLinkList(ctx context.Context) (homeLinkList []entity.HomeFriendLink, err error)
	}
)

var (
	localHome IHome
)

func Home() IHome {
	if localHome == nil {
		panic("implement not found for interface IHome, forgot register?")
	}
	return localHome
}

func RegisterHome(i IHome) {
	localHome = i
}
