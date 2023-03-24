package main

// go mod tidy
// go run ./cmd/main.go
// git init -> Example: go mod init github.com/LyoDekken/go-api

import (
	"github.com/LyoDekken/go-api/api/router"
)

func main() {
	//Inicializa o Router uilizando o Gin
	router.Initialize()
}
