package handler

import (
	"fmt"
	"net/http"

	"web-api/pkg"
	"web-api/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) PostUserHandler(c *gin.Context) {
	var userRequest pkg.User

	err := c.ShouldBindJSON(&userRequest)
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

	user, err := h.userService.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) GetUsersHandler(c *gin.Context) {

	users, err := h.userService.Read()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	var usersRes []pkg.UserResponse

	for _, u := range users {
		userRes := convertToUserRes(u)
		usersRes = append(usersRes, userRes)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usersRes,
	})
}

func (h *userHandler) FindByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.userService.FindByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func convertToUserRes(u pkg.User) pkg.UserResponse {
	return pkg.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Address:  u.Address,
	}
}
