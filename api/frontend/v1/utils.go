package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ResourceImageReq struct {
	g.Meta `path:"/api/front/resource/image" method:"post" tags:"资源模块" sm:"图片上传接口"`
	File   *ghttp.UploadFile `json:"file"`
}

type ResourceImageRes struct {
	Url string `json:"url"`
}
