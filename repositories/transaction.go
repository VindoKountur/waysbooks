package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindTransactionByUserID(ID int, status string) ([]models.Transaction, error)
	FindTransaction() ([]models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
	FindBooksByID(BooksID []int) ([]models.Book, error)
	FindUserTemporaryCart(UserID int) (string, error)
	EmptyUserCart(UserID int) (models.Profile, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) FindBooksByID(BooksID []int) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books, &BooksID).Error
	return books, err
}

func (r *repository) FindUserTemporaryCart(UserID int) (string, error) {
	var profile models.Profile
	err := r.db.Where("user_id = ?", UserID).First(&profile).Error
	return profile.CartTmp, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Where("id = ?", ID).Preload("Book").Preload("User").First(&transaction).Error
	return transaction, err
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Book").Find(&transactions).Error
	return transactions, err
}

func (r *repository) FindTransactionByUserID(UserID int, status string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	var err error
	if status == "" {
		err = r.db.Order("created_at desc").Preload("User").Preload("Book").Where("user_id = ?", UserID).Find(&transactions).Error
	} else {
		err = r.db.Order("created_at desc").Preload("User").Preload("Book").Where("user_id = ? AND status = ?", UserID, status).Find(&transactions).Error
	}
	return transactions, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.First(&transaction, orderId)

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *repository) EmptyUserCart(UserID int) (models.Profile, error) {
	var profile models.Profile
	r.db.Where("user_id = ?", UserID).First(&profile)
	profile.CartTmp = ""
	err := r.db.Save(&profile).Error
	return profile, err
}
