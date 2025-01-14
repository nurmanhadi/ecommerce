package unit

import (
	"ecommerce/pkg/infrastructure/security"
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := security.GenerateAccessToken("123")
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
func TestVerifyJwt(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzM0IiwiZXhwIjoxNzM2ODY3NzY2fQ.kdTM-32lzq4_s0qCZUs_O0K36ew42eOcfmD201BnaFo"
	userId, err := security.VerifyJwt(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(userId)
}
