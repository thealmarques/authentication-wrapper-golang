package interfaces

import "github.com/gin-gonic/gin"

// LoginService - Methods to perform login
type LoginService interface {
	Login(username string, password string) bool
}

// AuthenticationController - Login controller interface
type AuthenticationController interface {
	Login(ctx *gin.Context)
}
