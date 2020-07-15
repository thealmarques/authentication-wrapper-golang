package services

import (
	"authentication-layer/interfaces"
	"authentication-layer/models"
	"authentication-layer/utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtConfig struct {
	secretKey string
	issuer    string
}

// CreateJWTService - Returns an instance of JWT service
func CreateJWTService() interfaces.JWTService {
	return &jwtConfig{
		secretKey: utils.GetSecretKey(),
		issuer:    "alm-github",
	}
}

// GenerateToken - Generates JWT token
func (jwtService *jwtConfig) GenerateToken(username string, admin bool) string {
	// Set custom and standard claims
	claims := &models.JWTClaims{
		Username: username,
		Admin:    admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtService.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(jwtService.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// ValidateToken - validates token
func (jwtService *jwtConfig) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtService.secretKey), nil
	})
}
