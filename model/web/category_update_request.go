package web

type CategoryUpdateRequest struct {
	Name string `validate:"required,min=3,max=50" json:"name"`
}
