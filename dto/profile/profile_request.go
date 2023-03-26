package profiledto

type ProfileUpdateRequest struct {
	Phone   string `json:"phone" form:"phone"`
	Photo   string `json:"photo" form:"photo"`
	Gender  string `json:"gender" form:"gender"`
	Address string `json:"address" form:"address"`
}
