package repositories

import (
	"mgo-skeleton/bin/modules/auth/models"
	"mgo-skeleton/bin/pkg/helpers"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
	Register(req *models.UserModel) error
	GetUserByEmail(email string) (*models.UserModel, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) EmailExist(email string) bool {
	var user models.UserModel
	err := r.db.Table("users").Where("email = ?", email).First(&user).Error

	return err == nil
}

func (r *authRepository) Register(user *models.UserModel) error {
	// err := r.db.Create(&user).Error /// gorm: automatically added table name from struct, the name table is user_model
	err := r.db.Table("users").Create(&user).Error

	return err
}

func (r *authRepository) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel

	rows, err := r.db.Raw("select * from users where email = ? limit 1", email).Rows()
	if err != nil {
		return &user, err
	}

	defer rows.Close()

	if rows.Next() {
		r.db.ScanRows(rows, &user)
		return &user, nil
	}

	return &user, &helpers.NotFoundError{Message: "email not found"}
}
