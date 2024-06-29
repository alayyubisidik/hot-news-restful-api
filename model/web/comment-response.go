package web

import "time"

type CommentResponse struct {
	Id        int             `json:"id"`
	Content   string          `json:"content"`
	CreatedAt time.Time       `json:"created_at"`
	User      UserResponse    `json:"user"`
	Article   ArticleSimpleResponse `json:"article"`
}
