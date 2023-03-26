package models

import "time"

type Profile struct {
	ID            int          `json:"id" gorm:"primary_key:auto_increment"`
	Phone         string       `json:"phone" gorm:"varchar(255)"`
	Photo         string       `json:"photo" gorm:"varchar(255)"`
	PhotoPublicID string       `json:"photo_public_id" gorm:"varchar(255)"`
	Gender        string       `json:"gender" gorm:"varchar(255)"`
	Address       string       `json:"address" gorm:"varchar(255)"`
	UserID        int          `json:"user_id"`
	User          UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CartTmp       string       `json:"cart_tmp" gorm:"varchar(255)"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

// Associated with (User)
type ProfileResponse struct {
	UserID  int    `json:"-"`
	Phone   string `json:"phone"`
	Photo   string `json:"photo"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
