package profiledto

import "time"

type ProfileResponse struct {
	Phone     string    `json:"phone"`
	Photo     string    `json:"photo"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	UpdatedAt time.Time `json:"updated_at"`
}
