package repositories

import (
	"waysbooks/models"

	"gorm.io/gorm"
)

type BestBookResult struct {
	BookID int `json:"book_id"`
	Total  int `json:"total"`
}

type BookRepository interface {
	FindBook() ([]models.Book, error)
	FindBookByKeyword(keyword string) ([]models.Book, error)
	GetBook(ID int) (models.Book, error)
	CreateBook(Book models.Book) (models.Book, error)
	UpdateBook(Book models.Book) (models.Book, error)
	DeleteBook(Book models.Book) (models.Book, error)
	FindBestBook() ([]BestBookResult, error)
	CheckExistISBN(ISBN string) (models.Book, error)
}

func RepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBook() ([]models.Book, error) {
	var Books []models.Book
	err := r.db.Order("created_at desc").Find(&Books).Error
	return Books, err
}
func (r *repository) FindBookByKeyword(keyword string) ([]models.Book, error) {
	var Books []models.Book
	searchQuery := "%" + keyword + "%"
	err := r.db.Where("title LIKE ?", searchQuery).Or("author LIKE ?", searchQuery).Or("isbn LIKE ?", searchQuery).Find(&Books).Error
	return Books, err
}

func (r *repository) GetBook(ID int) (models.Book, error) {
	var Book models.Book
	err := r.db.First(&Book, ID).Error
	return Book, err
}

func (r *repository) CreateBook(Book models.Book) (models.Book, error) {
	err := r.db.Create(&Book).Error

	return Book, err
}

func (r *repository) UpdateBook(Book models.Book) (models.Book, error) {
	err := r.db.Save(&Book).Error
	return Book, err
}

func (r *repository) DeleteBook(Book models.Book) (models.Book, error) {
	err := r.db.Delete(&Book).Error
	return Book, err
}

func (r *repository) FindBestBook() ([]BestBookResult, error) {
	var bestBookResult []BestBookResult
	err := r.db.Table("cart").Select("count(book_id) as total", "book_id").Joins("inner join transactions transaction on cart.transaction_id = transaction.id").Where("transaction.status = ?", "success").Group("book_id").Order("total desc").Limit(3).Scan(&bestBookResult).Error
	return bestBookResult, err
}

func (r *repository) CheckExistISBN(ISBN string) (models.Book, error) {
	var book models.Book
	err := r.db.Where("isbn = ?", ISBN).First(&book).Error
	return book, err
}
