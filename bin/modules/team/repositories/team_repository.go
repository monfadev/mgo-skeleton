package repositories

import (
	"mgo-skeleton/bin/modules/team/models"

	"gorm.io/gorm"
)

type TeamRepository interface {
	Create(req *models.TeamModel) error
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *teamRepository {
	return &teamRepository{
		db: db,
	}
}

func (r *teamRepository) Create(user *models.TeamModel) error {
	err := r.db.Table("teams").Create(&user).Error

	return err
}
