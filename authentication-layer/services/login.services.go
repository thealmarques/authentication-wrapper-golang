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
	// TODO - connect to DB to check credentials
	return &loginService{
		authorizedUsername: "test",
		authorizedPassword: "password",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
