package auth

import (
	"fmt"
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

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)

	fmt.Println(err_token)

	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}

		ctx.JSON(http.StatusBadRequest, webResponse)

		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

