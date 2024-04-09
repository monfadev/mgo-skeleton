package services

import (
	"mgo-skeleton/bin/modules/team/models"
	"mgo-skeleton/bin/modules/team/repositories"
	"mgo-skeleton/bin/pkg/helpers"
)

type TeamService interface {
	Create(req *models.TeamRequest) error
}

type teamService struct {
	repositories repositories.TeamRepository
}

func NewTeamService(r repositories.TeamRepository) *teamService {
	return &teamService{
		repositories: r,
	}
}

func (s *teamService) Create(req *models.TeamRequest) error {

	if req.Password != req.PasswordConfirmation {
		return &helpers.BadRequestError{Message: "password not match"}
	}

	passwordHash, err := helpers.HashPassword(req.Password)
	if err != nil {
		return &helpers.InternalServerError{Message: err.Error()}
	}

	user := models.TeamModel{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: passwordHash,
		UserId:   req.UserId,
	}

	if err := s.repositories.Create(&user); err != nil {
		return &helpers.InternalServerError{Message: err.Error()}
	}

	return nil
}
