package controller

import (
	"net/http"

	"github.com/duongha/go-ecommerce/internal/service"
	"github.com/duongha/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.Success(c, http.StatusOK, "Get user by id")
}
