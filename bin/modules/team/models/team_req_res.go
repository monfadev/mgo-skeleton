package models

type TeamRequest struct {
	UserId               int    `json:"user_id"` ///foreign key with users.id
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Role                 string `json:"role" validate:"required,role"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirm" binding:"required"`
}

type TeamResponse struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
