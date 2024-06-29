package web

type CommentUpdateRequest struct {
	Id int `validate:"required" json:"id"`
	Content string `validate:"required,min=3,max=65535" json:"content"`
}