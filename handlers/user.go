package handlers

import (
	"net/http"
	"strconv"
	"time"
	dto "waysbooks/dto/result"
	usersdto "waysbooks/dto/users"
	"waysbooks/repositories"

	// "github.com/go-playground/validator"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: users})
}

func (h *handlerUser) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: user})
}

func (h *handlerUser) GetLoginUserInfo(c echo.Context) error {
	idUser := int(c.Get("userLogin").(jwt.MapClaims)["id"].(float64))
	user, err := h.UserRepository.GetUser(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: user})
}

func (h *handlerUser) UpdateLoginUser(c echo.Context) error {
	idUser := int(c.Get("userLogin").(jwt.MapClaims)["id"].(float64))
	user, err := h.UserRepository.GetUser(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	request := new(usersdto.UpdateUserRequest)
	if err = c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	user.UpdatedAt = time.Now()

	updatedUser, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: updatedUser})
}

// func convertResponseUser(u models.User, transactions []transactiondto.TransactionResponse) usersdto.UserTransactionByIDResponse {
// 	return usersdto.UserTransactionByIDResponse{
// 		ID:          u.ID,
// 		Name:        u.Name,
// 		Email:       u.Email,
// 		Profile:     u.Profile,
// 		Product:     u.Product,
// 		Addresses:   u.Addresses,
// 		Transaction: transactions,
// 		CreatedAt:   u.CreatedAt,
// 		UpdatedAt:   u.UpdatedAt,
// 	}
// }
