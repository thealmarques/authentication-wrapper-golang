package models

// Credentials user credentials
type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
