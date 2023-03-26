package cartdto

type CartResponse struct {
	BookCart   []int `json:"cart"`
	TotalPrice int   `json:"total_price"`
}
