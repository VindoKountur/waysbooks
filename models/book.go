package models

import "time"

type Book struct {
	ID                int       `json:"id" gorm:"primary_key:auto_increment"`
	Title             string    `json:"title" gorm:"type: varchar(255)"`
	Author            string    `json:"author" gorm:"type: varchar(255)"`
	PublicationDate   time.Time `json:"publication_date"`
	Pages             int       `json:"pages" gorm:"type: int"`
	Price             int       `json:"price" gorm:"type: int"`
	ISBN              string    `json:"isbn" gorm:"type: varchar(255)"`
	About             string    `json:"about" gorm:"type: text"`
	Thumbnail         string    `json:"thumbnail" gorm:"type: varchar(255)"`
	Content           string    `json:"content" gorm:"type: varchar(255)"`
	ThumbnailPublicID string    `json:"thumbnail_public_id" gorm:"varchar(255)"`
	ContentPublicID   string    `json:"content_public_id" gorm:"varchar(255)"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
