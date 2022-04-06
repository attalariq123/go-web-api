package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"web-api/pkg"
	"web-api/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService service.Service
}

func NewBookHandler(bookService service.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	var bookRequest pkg.Book

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errMessages = append(errMessages, errMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookRes := convertToBookRes(book)
	c.JSON(http.StatusOK, gin.H{
		"data": bookRes,
	})
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {

	books, err := h.bookService.Read()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": err,
		})
	}

	var booksRes []pkg.BookResponse

	for _, b := range books {
		bookRes := convertToBookRes(b)
		booksRes = append(booksRes, bookRes)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksRes,
	})
}

func (h *bookHandler) GetBookById(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	book, err := h.bookService.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	bookRes := convertToBookRes(book)

	c.JSON(http.StatusOK, gin.H{
		"data": bookRes,
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest pkg.Book

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errMessages = append(errMessages, errMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMessages,
		})
		return
	}

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookRes := convertToBookRes(book)

	c.JSON(http.StatusOK, gin.H{
		"data": bookRes,
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	book, err := h.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	bookRes := convertToBookRes(book)

	c.JSON(http.StatusOK, gin.H{
		"data": bookRes,
	})
}

func convertToBookRes(b pkg.Book) pkg.BookResponse {
	return pkg.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
