package models

import (
	"time"
)

type Transaction struct {
	ID         int          `json:"id"  gorm:"primary_key:auto_increment"`
	UserID     int          `json:"user_id" gorm:"type:int"`
	User       UserResponse `json:"user"`
	TotalPrice int          `json:"total_price"`
	Status     string       `json:"status" gorm:"type:varchar(255)"`
	Book       []Book       `json:"books" gorm:"many2many:cart;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BookID     []int        `json:"books_id" gorm:"-"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}
