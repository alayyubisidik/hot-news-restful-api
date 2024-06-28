package web

type ArticleCreateRequest struct {
	Title      string `validate:"required,min=3,max=255" json:"title"`
	Content    string `validate:"required,min=3,max=65535" json:"content"`
	UserId     int `validate:"required" json:"user_id"`
	CategoryId int `validate:"required" json:"category_id"`
}
