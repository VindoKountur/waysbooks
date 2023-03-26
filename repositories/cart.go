package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetTemporaryUserCart(UserID int) (models.Profile, error)
	UpdateTemporaryCart(profile models.Profile) (models.Profile, error)
	GetProductPrice(BookID int) (int, error)
	GetSuccessUserTransaction(UserID int) ([]models.Transaction, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetTemporaryUserCart(UserID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Where("user_id = ?", UserID).First(&profile).Error
	return profile, err
}

func (r *repository) UpdateTemporaryCart(profile models.Profile) (models.Profile, error) {
	err := r.db.Save(&profile).Error
	return profile, err
}

func (r *repository) GetProductPrice(BookID int) (int, error) {
	var book models.Book
	err := r.db.First(&book, BookID).Error
	return book.Price, err
}

func (r *repository) GetSuccessUserTransaction(UserID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Book").Where("user_id = ? && status = ?", UserID, "success").Find(&transactions).Error
	return transactions, err
}
