// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Literary_Snowlands/internal/model"
)

type (
	INews interface {
		FindLatestTwoNewsInfo() (newsInfo []model.NewsInfo, err error)
		FindNewsInfoById(id int) (newsInfo *model.NewsInfo, err error)
	}
)

var (
	localNews INews
)

func News() INews {
	if localNews == nil {
		panic("implement not found for interface INews, forgot register?")
	}
	return localNews
}

func RegisterNews(i INews) {
	localNews = i
}
