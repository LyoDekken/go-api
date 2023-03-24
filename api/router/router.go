package router

import (
	"log"
	"net/http"

	"time"

	"github.com/LyoDekken/go-api/config"
	"github.com/LyoDekken/go-api/helper"
	"github.com/LyoDekken/go-api/model"
	"github.com/LyoDekken/go-api/repositories"
	"github.com/go-playground/validator/v10"

	"github.com/LyoDekken/go-api/api/http/auth"
	"github.com/LyoDekken/go-api/api/http/controller"
	"github.com/LyoDekken/go-api/api/router/routes"
	"github.com/LyoDekken/go-api/api/service"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	database := config.ConnectionDB(&loadConfig)

	validate := validator.New()

	database.Table("users").AutoMigrate(&model.Users{})

	//Init Repository
	userRepository := repositories.NewUsersRepositoryImpl(database)

	//Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

	//Init controller
	auth := auth.NewAuthenticationController(authenticationService)
	controller := controller.NewAuthenticationController(authenticationService)

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,
			"Server listening on http://localhost:8080/api",
		)
	})

	router.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"code": "PAGE_NOT_FOUND", "message": "Page not found",
		})
	})

	handler := routes.InitializeRoutes(auth, controller, router)

	// Get the port from the environment
	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
