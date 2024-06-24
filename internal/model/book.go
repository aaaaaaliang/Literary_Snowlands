package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Comment struct {
	Id               uint64      `json:"id"`
	CommentContent   string      `json:"commentContent"`
	CommentUser      string      `json:"commentUser"`
	CommentUserId    uint64      `json:"commentUserId"`
	CommentUserPhoto string      `json:"commentUserPhoto"`
	CommentTime      *gtime.Time `json:"commentTime"`
}
