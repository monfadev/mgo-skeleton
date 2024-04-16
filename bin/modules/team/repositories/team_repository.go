package repositories

import (
	"errors"
	"fmt"
	"mgo-skeleton/bin/modules/team/models"
	"mgo-skeleton/bin/pkg/helpers"

	"gorm.io/gorm"
)

type TeamRepository interface {
	EmailExist(email string) bool
	Create(req *models.TeamModel) error

	TotalData(params *helpers.FilterParams, userId int) (int64, error)
	FindAll(params *helpers.FilterParams, userId int) (*[]models.TeamResponse, error)

	Detail(id int) (models.TeamResponse, error)
	Delete(id int, userId int) error
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *teamRepository {
	return &teamRepository{
		db: db,
	}
}

func (r *teamRepository) EmailExist(email string) bool {
	var user models.TeamModel
	err := r.db.Table("teams").Where("email = ?", email).First(&user).Error

	return err == nil
}

func (r *teamRepository) Create(user *models.TeamModel) error {
	err := r.db.Table("teams").Create(&user).Error

	return err
}

func (r *teamRepository) TotalData(params *helpers.FilterParams, userId int) (int64, error) {
	var response int64

	query := r.db.Debug().Table("teams").Where("user_id = ?", userId)

	if params.Search != "" {
		search := fmt.Sprintf("%%%s%%", params.Search)
		query.Where("lower(name) LIKE lower(?)", search)
	}

	err := query.Count(&response)
	if err != nil {
		return response, err.Error
	}

	return response, nil
}

func (r *teamRepository) FindAll(params *helpers.FilterParams, userId int) (*[]models.TeamResponse, error) {
	var response []models.TeamResponse

	query := r.db.Debug().Table("teams").Where("user_id = ?", userId)

	if params.Search != "" {
		search := fmt.Sprintf("%%%s%%", params.Search)
		query.Where("lower(name) LIKE lower(?)", search)
	}

	err := query.Offset(params.Offset).Limit(params.Limit).Find(&response).Error

	return &response, err
}

func (r *teamRepository) Detail(id int) (models.TeamResponse, error) {
	var response models.TeamResponse
	err := r.db.Table("teams").Where("id = ?", id).First(&response).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("user not found")
	}

	return response, err
}

func (r *teamRepository) Delete(id int, userId int) error {
	var req models.TeamRequest
	err := r.db.Debug().Table("teams").Where("id = ? and user_id = ?", id, userId).Delete(&req).Error
	fmt.Println("repository data is ", req)
	return err
}
