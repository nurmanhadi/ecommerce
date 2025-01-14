package security

import (
	"ecommerce/config"
	"ecommerce/pkg/infrastructure/exception"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claim struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId string) (string, error) {
	exp := time.Now().Add((time.Hour * 24) * time.Duration(config.Viper.Jwt.ExpiresAt))
	claims := claim{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(config.Viper.Jwt.SecretKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}
func VerifyJwt(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Viper.Jwt.SecretKey), nil
	})
	if err != nil {
		return "", exception.JwtInvalidSignature
	}
	claims, ok := token.Claims.(*claim)
	if !ok {
		return "", errors.New("undefined token")
	}
	if claims.ExpiresAt.Before(time.Now()) {
		return "", exception.JwtExpired
	}
	return claims.Id, nil
}
