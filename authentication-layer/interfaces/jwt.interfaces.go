package interfaces

import "github.com/dgrijalva/jwt-go"

// JWTService interface
type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}
