package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	profiledto "waysbooks/dto/profile"
	dto "waysbooks/dto/result"
	"waysbooks/models"
	"waysbooks/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) UpdateProfileByUser(c echo.Context) error {
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	userLogin := c.Get("userLogin")
	idUserLogin := userLogin.(jwt.MapClaims)["id"].(float64)

	dataFilePhoto := c.Get("dataFileProfile").(string)

	profile, err := h.ProfileRepository.GetProfileByUser(int(idUserLogin))
	request := profiledto.ProfileUpdateRequest{
		Photo:   dataFilePhoto,
		Phone:   c.FormValue("phone"),
		Gender:  c.FormValue("gender"),
		Address: c.FormValue("address"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	if request.Phone != "" {
		profile.Phone = request.Phone
	}
	if request.Gender != "" {
		profile.Gender = request.Gender
	}
	if request.Address != "" {
		profile.Address = request.Address
	}
	if dataFilePhoto != "" {
		if profile.PhotoPublicID != "" {
			_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: profile.PhotoPublicID})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		resp, err := cld.Upload.Upload(ctx, request.Photo, uploader.UploadParams{Folder: "waysbooks/profile"})

		if err != nil {
			fmt.Println(err.Error())
		}
		profile.Photo = resp.SecureURL
		profile.PhotoPublicID = resp.PublicID
	}
	profile.UpdatedAt = time.Now()

	updatedProfile, err := h.ProfileRepository.UpdateProfileByUser(profile, profile.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertProfileResponse(updatedProfile)})
}

func convertProfileResponse(profile models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		Phone:     profile.Phone,
		Photo:     profile.Photo,
		Gender:    profile.Gender,
		Address:   profile.Address,
		UpdatedAt: profile.UpdatedAt,
	}
}
