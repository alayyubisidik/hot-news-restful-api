package web

type LikeCreateRequest struct {
	UserId int `validate:"required" json:"user_id"`
	ArticleId int `validate:"required" json:"article_id"`
}