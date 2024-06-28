package web

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Username string `validate:"required,min=3,max=100" json:"username"`
	FullName string `validate:"required,min=3,max=100" json:"full_name"`
	Email    string `validate:"required,email" json:"email"`
}
