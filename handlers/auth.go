package handlers

import (
	"log"
	"net/http"
	"time"
	authdto "waysbooks/dto/auth"
	dto "waysbooks/dto/result"
	"waysbooks/models"
	"waysbooks/pkg/bcrypt"
	jwtToken "waysbooks/pkg/jwt"
	"waysbooks/repositories"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository,

) *handlerAuth {
	return &handlerAuth{
		AuthRepository,
	}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// Cek if email is exist
	checkUserMail, err := h.AuthRepository.Login(request.Email)
	if err != nil && (err.Error() != "record not found") {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Error checking email" + err.Error()})
	}

	if checkUserMail.ID != 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email already exists"})
	}

	hashedPassword, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var role = "user"
	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashedPassword,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// User login
	claims := jwt.MapClaims{}
	claims["id"] = data.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Print(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginReponse{
		Email: user.Email,
		Role:  user.Role,
		Token: token,
		Photo: "",
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: loginResponse})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Wrong email or password"})
		} else {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Wrong email or password"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Print(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginReponse{
		Email: user.Email,
		Role:  user.Role,
		Token: token,
		Photo: user.Profile.Photo,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: loginResponse})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	idUserLogin := int(c.Get("userLogin").(jwt.MapClaims)["id"].(float64))
	user, err := h.AuthRepository.CheckAuth(idUserLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	response := authdto.LoginReponse{
		Email: user.Email,
		Role:  user.Role,
		Photo: user.Profile.Photo,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: response})
}
