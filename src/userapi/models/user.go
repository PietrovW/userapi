package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required,min=5,max=20"`
	Email string `json:"email" validate:"required,email"`
}
