package web

import (
	"time"
)

type ArticleResponse struct {
	Id        int              `json:"id"`
	Title     string           `json:"title"`
	Slug      string           `json:"slug"`
	Content   string           `json:"content"`
	CreatedAt time.Time        `json:"created_at"`
	User      UserResponse     `json:"user"`
	Category  CategoryResponse `json:"category"`
}

type ArticleSimpleResponse struct {
	Id        int              `json:"id"`
	Title     string           `json:"title"`
	Slug      string           `json:"slug"`
	Content   string           `json:"content"`
	CreatedAt time.Time        `json:"created_at"`
}
