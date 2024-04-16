package models

type TeamRequest struct {
	UserId               int    `json:"user_id"` ///foreign key with users.id
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Role                 string `json:"role" validate:"required"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirm"`
}

type TeamResponse struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
