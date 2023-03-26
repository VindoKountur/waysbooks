package transactiondto

type CreateTransactionRequest struct {
	UserID     int    `json:"user_id" validate:"required"`
	TotalPrice int    `json:"total_price" validate:"required"`
	BookID     []int  `json:"books_id" validate:"required"`
	Status     string `json:"status" validate:"required"`
}
