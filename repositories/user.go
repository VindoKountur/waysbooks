package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (h *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := h.db.Preload("Profile").Find(&users).Error
	return users, err
}

func (h *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := h.db.Preload("Profile").First(&user, ID).Error
	return user, err
}

func (h *repository) CreateUser(user models.User) (models.User, error) {
	err := h.db.Create(&user).Error
	return user, err
}

func (h *repository) UpdateUser(user models.User) (models.User, error) {
	err := h.db.Save(&user).Error
	return user, err
}

func (h *repository) DeleteUser(user models.User) (models.User, error) {
	err := h.db.Delete(&user).Error
	return user, err
}
