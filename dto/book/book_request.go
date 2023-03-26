package bookdto

import "time"

type CreateBookRequest struct {
	Title           string    `json:"title" from:"title" validate:"required"`
	Author          string    `json:"author" from:"author" validate:"required"`
	PublicationDate time.Time `json:"publication_date" from:"publication_date" validate:"required"`
	Pages           int       `json:"pages" from:"pages" validate:"required"`
	ISBN            string    `json:"isbn" from:"isbn" validate:"required"`
	Price           int       `json:"price" from:"price" validate:"required"`
	About           string    `json:"about" from:"about" validate:"required"`
	Thumbnail       string    `json:"thumbnail" from:"thumbnail" validate:"required"`
	Content         string    `json:"content" from:"content" validate:"required"`
}
type UpdateBookRequest struct {
	Title           string    `json:"title" from:"title"`
	Author          string    `json:"author" from:"author"`
	PublicationDate time.Time `json:"publication_date" from:"publication_date"`
	Pages           int       `json:"pages" from:"pages"`
	ISBN            string    `json:"isbn" from:"isbn"`
	Price           int       `json:"price" from:"price"`
	About           string    `json:"about" from:"about"`
	Thumbnail       string    `json:"thumbnail" from:"thumbnail"`
	Content         string    `json:"content" from:"content"`
}
