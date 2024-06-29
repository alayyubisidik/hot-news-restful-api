package web

import (
	"time"
)

type LikeResponse struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	User      UserResponse
	Article   ArticleSimpleResponse
}
