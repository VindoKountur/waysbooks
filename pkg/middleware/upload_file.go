package middleware

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	dto "waysbooks/dto/result"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const imgExt = "png"
const pdfExt = "pdf"

func UploadPhotoProfile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("photo")
		method := c.Request().Method

		if err != nil {
			if method == "PATCH" && err.Error() == "http: no such file" {
				c.Set("dataFileProfile", "")
				return next(c)
			}
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		data, err := handleFile(file, imgExt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		c.Set("dataFileProfile", data)
		return next(c)
	}
}
func UploadThumbnail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("thumbnail")
		method := c.Request().Method

		if err != nil {
			if method == "PATCH" && err.Error() == "http: no such file" {
				c.Set("dataFileThumbnail", "")
				return next(c)
			}
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		data, err := handleFile(file, imgExt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		c.Set("dataFileThumbnail", data)
		return next(c)
	}
}
func UploadPDF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("content")
		method := c.Request().Method

		if err != nil {
			if method == "PATCH" && err.Error() == "http: no such file" {
				c.Set("dataFilePDF", "")
				return next(c)
			}
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		// Setup S3 Upload
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Printf("error: %v", err)
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Error setup S3" + err.Error()})
		}
		client := s3.NewFromConfig(cfg)

		// Open file
		openedFile, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Error setup S3" + err.Error()})
		}

		uploader := manager.NewUploader(client)
		result, errUpload := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("waysbooks"),
			Key:    aws.String(uuid.NewString() + ".pdf"),
			Body:   openedFile,
			ACL:    "public-read",
		})
		if errUpload != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Error upload S3" + err.Error()})
		}

		// data, err := handleFile(file, pdfExt)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		// }

		// c.Set("dataFilePDF", result.Location)
		c.Set("dataFilePDF", result.Location)
		return next(c)
	}
}

func handleFile(file *multipart.FileHeader, ext string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	tempFile, err := ioutil.TempFile("uploads", "file-*."+ext)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	if _, err = io.Copy(tempFile, src); err != nil {
		return "", err
	}

	data := tempFile.Name()
	return data, nil
}
