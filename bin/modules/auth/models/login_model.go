package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
