package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	var profile models.Profile
	profile.UserID = user.ID
	err = r.db.Create(&profile).Error
	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email =?", email).Preload("Profile").First(&user).Error
	return user, err
}

func (r *repository) CheckAuth(ID int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", ID).Preload("Profile").First(&user).Error
	return user, err
}
