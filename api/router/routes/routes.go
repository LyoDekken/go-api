package routes

import (
	"github.com/LyoDekken/go-api/api/http/controller"
	"github.com/LyoDekken/go-api/api/http/auth"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(
	auth *auth.AuthenticationController,
	controller *controller.AuthenticationController, 
	router *gin.Engine,
	) *gin.Engine {

	basePath := "/api"


	api := router.Group(basePath)
	{
		api.POST("/login", auth.Login)
		api.POST("/register", controller.Register)
	}

	return router
}
