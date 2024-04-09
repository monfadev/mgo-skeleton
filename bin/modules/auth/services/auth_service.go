package services

import (
	"mgo-skeleton/bin/configs"
	"mgo-skeleton/bin/modules/auth/models"
	"mgo-skeleton/bin/modules/auth/repositories"
	"mgo-skeleton/bin/pkg/helpers"
)

type AuthServices interface {
	Register(req *models.RegisterRequest) error
	Login(req *models.LoginRequest) (*models.LoginResponse, error)
}

type authService struct {
	repositories repositories.AuthRepository
}

func NewAuthServices(r repositories.AuthRepository) *authService {
	return &authService{
		repositories: r,
	}
}

func (s *authService) Register(req *models.RegisterRequest) error {
	if emailExist := s.repositories.EmailExist(req.Email); emailExist {
		return &helpers.BadRequestError{Message: "email already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return &helpers.BadRequestError{Message: "password not match"}
	}

	passwordHash, err := helpers.HashPassword(req.Password)
	if err != nil {
		return &helpers.InternalServerError{Message: err.Error()}
	}

	user := models.UserModel{
		Name:     req.Name,
		Email:    req.Email,
		Role:     configs.ROLE_ADMIN,
		Password: passwordHash,
	}

	if err := s.repositories.Register(&user); err != nil {
		return &helpers.InternalServerError{Message: err.Error()}
	}

	return nil

}

func (s *authService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	var response models.LoginResponse

	user, err := s.repositories.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &helpers.NotFoundError{Message: "wrong email"}
	}

	if err := helpers.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &helpers.NotFoundError{Message: "wrong password"}
	}

	token, err := helpers.GenerateToken(user)

	if err != nil {
		return nil, &helpers.InternalServerError{Message: err.Error()}
	}

	response = models.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Role:  user.Role,
		Token: token,
	}

	return &response, nil

}
