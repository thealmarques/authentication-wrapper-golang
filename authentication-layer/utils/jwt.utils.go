package utils

import "os"

// GetSecretKey - gets secret key from environment variables if exist
func GetSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
