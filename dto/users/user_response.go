package usersdto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UserTransactionResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
