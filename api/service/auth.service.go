package service

import (
	"errors"

	"github.com/LyoDekken/go-api/api/data/request"
	"github.com/LyoDekken/go-api/config"
	"github.com/LyoDekken/go-api/function"
	"github.com/LyoDekken/go-api/helper"
	"github.com/LyoDekken/go-api/model"
	"github.com/LyoDekken/go-api/repositories"

	"github.com/go-playground/validator/v10"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUsersRequest)
}

type AuthenticationServiceImpl struct {
	UsersRepository repositories.UsersRepository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repositories.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in database
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := function.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := function.GenerateToken(config.TokenExpiresIn, new_users.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(users request.CreateUsersRequest) {

	hashedPassword, err := function.HashPassword(users.Password)

	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UsersRepository.Save(newUser)
}
