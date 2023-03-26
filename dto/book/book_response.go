package bookdto

import "time"

type BookResponse struct {
	ID              int       `json:"id" validate:"required"`
	Title           string    `json:"title" validate:"required"`
	Author          string    `json:"author" validate:"required"`
	PublicationDate time.Time `json:"publication_date" validate:"required"`
	Pages           int       `json:"pages" validate:"required"`
	ISBN            string    `json:"isbn" validate:"required"`
	Price           int       `json:"price" validate:"required"`
	About           string    `json:"about" validate:"required"`
	Thumbnail       string    `json:"thumbnail" validate:"required"`
	Content         string    `json:"content" validate:"required"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
	UpdatedAt       time.Time `json:"updated_at" validate:"required"`
}
