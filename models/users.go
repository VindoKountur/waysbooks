package models

import "time"

type User struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"-" gorm:"type: varchar(255)"`
	Role      string          `json:"role" gorm:"type: varchar(255)"`
	Name      string          `json:"name" gorm:"type: varchar(255)"`
	Profile   ProfileResponse `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	// Transaction []TransactionResponse `json:"transactions"`
	// Product     []ProductUserResponse `json:"products"`
	// Addresses   []AddressesResponse   `json:"addresses"`
}

// Associated with (Transaction, Profile, Address)
type UserResponse struct {
	ID    int    `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UserResponse) TableName() string {
	return "users"
}
