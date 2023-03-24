package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/LyoDekken/go-api/config"
	"github.com/LyoDekken/go-api/repositories"
	"github.com/LyoDekken/go-api/function"

	"github.com/gin-gonic/gin"
)

// Authorize é um middleware que valida e decodifica o token de autenticação do usuário
func Authorize(userRepository repositories.UsersRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var token string

		// Verifica o Header da Autenticação
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Authorization header not provided",
			})
			return
		}

		// Verifica se o token está no formato correto
		fields := strings.Fields(authHeader)
		if len(fields) != 2 || fields[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Invalid Authorization header format. Use 'Bearer <token>'",
			})
			return
		}

		// Extrai o token do header da autenticação
		token = fields[1]

		// Valida o token
		cfg, err := config.LoadConfig(".")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "fail",
				"message": "Failed to load server configuration",
			})
			return
		}

		sub, err := function.ValidateToken(token, cfg.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Invalid or expired token",
			})
			return
		}

		// Busca o usuário no banco de dados pelo ID contido no token
		id, err := strconv.Atoi(fmt.Sprint(sub))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "fail",
				"message": "Failed to extract user ID from token",
			})
			return
		}

		user, err := userRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "fail",
				"message": "User associated with the token no longer exists",
			})
			return
		}

		// Adiciona o nome do usuário ao contexto da requisição
		ctx.Set("currentUser", user.Username)
		ctx.Next()
	}
}
