package handler

import (
	"net/http"
	"crowdfunding-rest-api/user"
	"crowdfunding-rest-api/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		
		// message error
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		
		// message error
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	// message success
	c.JSON(http.StatusOK, response)
}