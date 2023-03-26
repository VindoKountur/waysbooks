package handlers

import (
	"net/http"
	"strconv"
	"strings"
	cartdto "waysbooks/dto/cart"
	dto "waysbooks/dto/result"
	"waysbooks/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) AddToCart(c echo.Context) error {
	userLogin := c.Get("userLogin")
	idUserLogin := int(userLogin.(jwt.MapClaims)["id"].(float64))

	bookId := c.Param("id")
	bookIdInt, _ := strconv.Atoi(bookId)

	// Cek if user already purchased this book
	userTransactions, err := h.CartRepository.GetSuccessUserTransaction(idUserLogin)
	var idBooksPurchased []int
	for _, transaction := range userTransactions {
		for _, book := range transaction.Book {
			idBooksPurchased = append(idBooksPurchased, book.ID)
		}
	}

	var bookIsPurchased bool
	for _, purchasedBookId := range idBooksPurchased {
		if purchasedBookId == bookIdInt {
			bookIsPurchased = true
		}
	}


	if bookIsPurchased {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Book already purchased"})
	}

	// Manipulate user cart
	profile, err := h.CartRepository.GetTemporaryUserCart(idUserLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	userCart := strings.Split(profile.CartTmp, ",")
	var separator string
	if profile.CartTmp == "" {
		userCart = append(userCart, bookId)
		separator = ""
	} else {
		for _, item := range userCart {
			if item == bookId {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Book already in cart"})
			}
		}
		userCart = append(userCart, bookId)
		separator = ","
	}
	arrCart := strings.Join(userCart, separator)

	profile.CartTmp = arrCart
	_, err = h.CartRepository.UpdateTemporaryCart(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed to update cart"})
	}

	_, err = h.CartRepository.GetTemporaryUserCart(idUserLogin)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: "Book added to cart"})
}

func (h *handlerCart) RemoveBookFromCart(c echo.Context) error {

	userLogin := c.Get("userLogin")
	idUserLogin := int(userLogin.(jwt.MapClaims)["id"].(float64))

	bookId := c.Param("id")

	profile, err := h.CartRepository.GetTemporaryUserCart(idUserLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	userCart := strings.Split(profile.CartTmp, ",")
	findIdx := indexOf(bookId, userCart)

	if findIdx == -1 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Book is not in cart"})
	}
	filteredCart := remove(userCart, findIdx)
	var separator string
	if len(filteredCart) == 0 {
		separator = ""
	} else {
		separator = ","
	}

	newCart := strings.Join(filteredCart, separator)

	profile.CartTmp = newCart
	_, err = h.CartRepository.UpdateTemporaryCart(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: "Book removed from cart"})
}

func (h *handlerCart) GetUserCartList(c echo.Context) error {
	idUserLogin := int((c.Get("userLogin").(jwt.MapClaims)["id"]).(float64))

	userProfile, err := h.CartRepository.GetTemporaryUserCart(idUserLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	userCart := userProfile.CartTmp
	var cartResp cartdto.CartResponse
	if userCart == "" {
		cartResp = cartdto.CartResponse{
			BookCart:   []int{},
			TotalPrice: 0,
		}
		return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: cartResp})
	}

	arrUserCart := strings.Split(userCart, ",")

	for _, item := range arrUserCart {
		itemInt, _ := strconv.Atoi(item)
		cartResp.BookCart = append(cartResp.BookCart, itemInt)

		// Get Book Price
		price, err := h.CartRepository.GetProductPrice(itemInt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		}

		cartResp.TotalPrice += price
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: cartResp})
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
