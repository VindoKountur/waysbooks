package authdto

type LoginReponse struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
	Photo string `json:"photo"`
}
