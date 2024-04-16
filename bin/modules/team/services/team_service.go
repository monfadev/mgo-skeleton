package services

import (
	"fmt"
	"math"
	"mgo-skeleton/bin/modules/team/models"
	"mgo-skeleton/bin/modules/team/repositories"
	"mgo-skeleton/bin/pkg/helpers"
	"time"

	"github.com/go-playground/validator/v10"
)

type TeamService interface {
	Create(req *models.TeamRequest) error
	FindAll(params *helpers.FilterParams, userId int) (*[]models.TeamResponse, *helpers.Paginate, error)
	Detail(id int) (*models.TeamResponse, error)
	Delete(id int, userId int) error
	Update(id int, req *models.TeamRequest) error
}

type teamService struct {
	repositories repositories.TeamRepository
	validator    *validator.Validate
}

func NewTeamService(r repositories.TeamRepository) *teamService {
	return &teamService{
		repositories: r,
		validator:    validator.New(),
	}
}

func (s *teamService) Create(req *models.TeamRequest) error {

	if emailExist := s.repositories.EmailExist(req.Email); emailExist {
		return &helpers.BadRequestError{Message: "email already registered", MessageDev: "email already registered in teams"}
	}

	if req.Password != req.PasswordConfirmation || req.PasswordConfirmation != req.Password {
		return &helpers.BadRequestError{Message: "password not match", MessageDev: "password is not match in teams"}
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

func (s *teamService) FindAll(params *helpers.FilterParams, userId int) (*[]models.TeamResponse, *helpers.Paginate, error) {
	fmt.Println("params limit is ", params.Limit)
	count, err := s.repositories.TotalData(params, userId)
	fmt.Println("count is ", count)
	if err != nil {
		return nil, nil, &helpers.InternalServerError{Message: err.Error()}
	}

	response, err := s.repositories.FindAll(params, userId)
	if err != nil {
		return nil, nil, &helpers.InternalServerError{Message: err.Error()}
	}

	paginate := &helpers.Paginate{
		Total:     int(count),
		PerPage:   params.Limit,
		Page:      params.Page,
		TotalPage: int(math.Ceil(float64(count) / float64(params.Limit))),
	}

	return response, paginate, nil

}

func (s *teamService) Detail(id int) (*models.TeamResponse, error) {
	response, err := s.repositories.Detail(id)

	if err != nil {
		return nil, &helpers.BadRequestError{Message: "user not found"}
	}

	return &response, nil
}

func (s *teamService) Delete(id int, userId int) error {
	detailUser, err := s.Detail(id)
	if err != nil {
		return err
	}

	fmt.Println("delete team id is ", id)
	fmt.Println("delete team userId is ", userId)
	fmt.Println("delete team detailUser.UserId is ", detailUser.UserId)

	if userId != detailUser.UserId {
		return &helpers.BadRequestError{Message: "it is not your user team"}
	}

	if err := s.repositories.Delete(id, userId); err != nil {
		return &helpers.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *teamService) Update(id int, req *models.TeamRequest) error {
	fmt.Println("update service is ", req)

	if err := s.validator.Struct(req); err != nil {
		return &helpers.BadRequestError{Message: err.Error(), MessageDev: "struct is not match with team request"}
	}

	detailUser, err := (s).Detail(id)

	fmt.Println("detailUser is ", detailUser)

	if err != nil {
		return err
	}

	if req.UserId != detailUser.UserId {
		return &helpers.BadRequestError{Message: "it is not your user team", MessageDev: "differences fk user_id"}
	}

	user := models.TeamModel{
		Name:      req.Name,
		Email:     req.Email,
		Role:      req.Role,
		UpdatedAt: time.Now(),
	}

	if err := s.repositories.Update(id, &user); err != nil {
		return &helpers.InternalServerError{Message: err.Error()}
	}

	return err
}
