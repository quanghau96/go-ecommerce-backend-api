package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quanghau96/go-ecommerce-backend-api/internal/service"
	"github.com/quanghau96/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc UserController) GetUserById(c *gin.Context) {
	// data := uc.userService.GetInfoUser()
	// response.SuccessReponse(c, 20001, data)
	response.ErrorResponse(c, 20003)
}
