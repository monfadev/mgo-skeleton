package models

type TeamRequest struct {
	UserId               int    `json:"user_id"` ///foreign key with users.id
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Role                 string `json:"role" validate:"required,role"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirm"`
}

type TeamResponse struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
