package frontend

import (
	"Literary_Snowlands/api/frontend"
	v1 "Literary_Snowlands/api/frontend/v1"
	"Literary_Snowlands/utility"
	"context"
)

type UtilControllerV1 struct {
}

func NewUtilControllerV1() frontend.IUtilV1 {
	return &UtilControllerV1{}
}

func (c *UtilControllerV1) ResourceImages(ctx context.Context, req *v1.ResourceImageReq) (res *v1.ResourceImageRes, err error) {
	url, err := utility.ResourceImages(ctx, req.File)
	if err != nil {
		return nil, err
	}
	urlRes := &v1.ResourceImageRes{Url: url}
	return urlRes, nil
}
