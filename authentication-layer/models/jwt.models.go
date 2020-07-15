package models

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims model structure for jwt claims
type JWTClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}
