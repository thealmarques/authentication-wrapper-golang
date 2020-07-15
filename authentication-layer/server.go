package main

import (
	"authentication-layer/controllers"
	"authentication-layer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	jwtService   = services.CreateJWTService()
	loginService = services.CreateLoginService()

	authenticationController = controllers.CreateAuthenticationController(loginService, jwtService)
)

func main() {
	router := gin.Default()

	router.POST("/login", func(ctx *gin.Context) {
		token := authenticationController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not valid"})
		}
	})

	router.Run("0.0.0.0:9000")
}
