package controllers

import (
	"authentication-layer/interfaces"
	"authentication-layer/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authenticationController struct {
	loginService interfaces.LoginService
	jWtService   interfaces.JWTService
}

// CreateAuthenticationController - Creates authentication controller
func CreateAuthenticationController(loginService interfaces.LoginService,
	jWtService interfaces.JWTService) interfaces.AuthenticationController {
	return &authenticationController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *authenticationController) Login(ctx *gin.Context) {
	var credentials models.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong credentials"})
		return
	}

	path := ctx.Param("login")
	userKey := strings.Split(path, "/")[1]
	if userKey == "login" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User key not specified"})
		return
	}

	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		token := controller.jWtService.GenerateToken(credentials.Username, true)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	}
}
