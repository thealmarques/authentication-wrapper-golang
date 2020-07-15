package services

import (
	"authentication-layer/interfaces"
)

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

// CreateLoginService - Creates login service (static usernames for now)
func CreateLoginService() interfaces.LoginService {
	return &loginService{
		authorizedUsername: "test",
		authorizedPassword: "password",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
