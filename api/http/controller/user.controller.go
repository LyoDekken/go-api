package controller

import (
	"net/http"

	"github.com/LyoDekken/go-api/api/data/request"
	"github.com/LyoDekken/go-api/api/data/response"
	"github.com/LyoDekken/go-api/api/service"
	"github.com/LyoDekken/go-api/helper"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	controller.authenticationService.Register(createUsersRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
