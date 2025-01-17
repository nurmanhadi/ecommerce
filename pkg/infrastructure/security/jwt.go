package security

import (
	"ecommerce/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claim struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId string) (string, error) {
	exp := time.Now().Add((time.Hour * 24) * time.Duration(config.Viper.Jwt.Exp))
	claims := claim{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(config.Viper.Jwt.Key))
	if err != nil {
		return "", err
	}
	return ss, nil
}
func VerifyJwt(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Viper.Jwt.Key), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*claim)
	if !ok {
		return "", errors.New("undefined token")
	}
	return claims.Id, nil
}
