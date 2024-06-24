package frontend

import (
	"Literary_Snowlands/api/frontend"
	v1 "Literary_Snowlands/api/frontend/v1"
	"Literary_Snowlands/internal/service"
	"context"
)

type NewsControllerV1 struct {
}

func NewNewsControllerV1() frontend.INewsV1 {
	return &NewsControllerV1{}
}

func (c *NewsControllerV1) QueryNewsLatestInfo(ctx context.Context, req *v1.QueryNewsLatestInfoReq) (res []*v1.QueryNewsLatestInfoRes, err error) {
	latestNewsInfo, err := service.News().FindLatestTwoNewsInfo()
	if err != nil {
		return nil, err
	}
	var latestNewsArrayInfoRes []*v1.QueryNewsLatestInfoRes
	for _, v := range latestNewsInfo {
		infoRes := &v1.QueryNewsLatestInfoRes{
			NewsInfo: v1.NewsInfo{
				Id:           v.Id,
				CategoryId:   v.CategoryId,
				CategoryName: v.CategoryName,
				SourceName:   v.SourceName,
				Title:        v.Title,
				UpdateTime:   v.UpdateTime,
				Content:      v.Content,
			},
		}
		latestNewsArrayInfoRes = append(latestNewsArrayInfoRes, infoRes)
	}
	return latestNewsArrayInfoRes, err
}

func (c *NewsControllerV1) QueryNewsInfo(ctx context.Context, req *v1.QueryNewsInfoReq) (res *v1.QueryNewsInfoRes, err error) {
	newsInfo, err := service.News().FindNewsInfoById(req.Id)
	if err != nil {
		return nil, err
	}
	newsInfoRes := v1.QueryNewsInfoRes{NewsInfo: v1.NewsInfo{
		Id:           newsInfo.Id,
		CategoryId:   newsInfo.CategoryId,
		CategoryName: newsInfo.CategoryName,
		SourceName:   newsInfo.SourceName,
		Title:        newsInfo.Title,
		UpdateTime:   newsInfo.UpdateTime,
		Content:      newsInfo.Content,
	}}
	return &newsInfoRes, err
}
