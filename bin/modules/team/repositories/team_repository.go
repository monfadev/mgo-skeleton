package repositories

import (
	"errors"
	"fmt"
	"mgo-skeleton/bin/modules/team/models"

	"gorm.io/gorm"
)

type TeamRepository interface {
	EmailExist(email string) bool
	Create(req *models.TeamModel) error
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
