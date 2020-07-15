package main

import (
	"authentication-layer/controllers"
	"authentication-layer/middlewares"
	"authentication-layer/services"

	"github.com/gin-gonic/gin"
)

var (
	jwtService   = services.CreateJWTService()
	loginService = services.CreateLoginService()

	authenticationController = controllers.CreateAuthenticationController(loginService, jwtService)
)

func main() {
	router := gin.Default()

	router.POST("/*login", authenticationController.Login)

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, "Token is valid")
		})
	}

	router.Run("0.0.0.0:9000")
}
