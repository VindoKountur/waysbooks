package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfileByUser(userID int) (models.Profile, error)
	UpdateProfileByUser(profile models.Profile, UserID int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetProfileByUser(userID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Where("user_id =?", userID).First(&profile).Error
	return profile, err
}

func (r *repository) UpdateProfileByUser(profile models.Profile, UserID int) (models.Profile, error) {
	err := r.db.Where("user_id = ?", UserID).Save(&profile).Error
	return profile, err
}
