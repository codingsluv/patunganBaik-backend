package handler

import (
	"net/http"

	"github.com/codingsluv/crowdfounding/helper"
	"github.com/codingsluv/crowdfounding/user"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := helper.ApiResponse("User Registered Successfully", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)
}
