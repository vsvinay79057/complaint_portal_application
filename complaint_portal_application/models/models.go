package models

type RegisterRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	SecretCode string `json:"secret_code" validate:"required"`
}
type ComplaintRequest struct {
	Title   string `json:"title" validate:"required"`
	Summary string `json:"summary" validate:"required"`
	Rating  int    `json:"rating" validate:"required,min=1,max=5"`
}
