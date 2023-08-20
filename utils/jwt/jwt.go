package jwt

import (
	"thresher/utils/config"

	"github.com/golang-jwt/jwt"
)

func ValidateJwt(token string) (*jwt.Token, error) {
	jwtSecret := config.LoadConfig().Token.JwtSecret
	return jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
}