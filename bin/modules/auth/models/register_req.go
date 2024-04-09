package models

type RegisterRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Role                 string `json:"role" validate:"required,role"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirm"`
}
