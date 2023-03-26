package transactiondto

import (
	"time"
	bookdto "waysbooks/dto/book"
	usersdto "waysbooks/dto/users"
)

type TransactionResponse struct {
	ID         int                              `json:"id"`
	User       usersdto.UserTransactionResponse `json:"user"`
	TotalPrice int                              `json:"total_price"`
	Status     string                           `json:"status"`
	Books      []bookdto.BookResponse           `json:"books"`
	CreatedAt  time.Time                        `json:"created_at"`
	UpdatedAt  time.Time                        `json:"updated_at"`
}
