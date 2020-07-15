package interfaces

import "github.com/gin-gonic/gin"

// LoginService - Methods to perform login
type LoginService interface {
	Login(username string, password string) bool
}

// LoginController - Login controller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}
