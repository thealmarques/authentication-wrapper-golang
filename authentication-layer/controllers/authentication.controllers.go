package controllers

import (
	"authentication-layer/interfaces"
	"authentication-layer/models"

	"github.com/gin-gonic/gin"
)

type authenticationController struct {
	loginService interfaces.LoginService
	jWtService   interfaces.JWTService
}

// CreateAuthenticationController - Creates authentication controller
func CreateAuthenticationController(loginService interfaces.LoginService,
	jWtService interfaces.JWTService) interfaces.LoginController {
	return &authenticationController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *authenticationController) Login(ctx *gin.Context) string {
	var credentials models.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
