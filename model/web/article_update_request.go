package web

type ArticleUpdateRequest struct {
	Title      string `validate:"required,min=3,max=255" json:"title"`
	Content    string `validate:"required,min=3,max=65535" json:"content"`
	CategoryId int `validate:"required" json:"category_id"`
}
