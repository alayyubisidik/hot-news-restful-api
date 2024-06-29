package web

type CommentCreateRequest struct {
	UserId int `validate:"required" json:"user_id"`
	ArticleId int `validate:"required" json:"article_id"`
	Content string `validate:"required,min=3,max=65535" json:"content"`
}