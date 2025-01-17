package unit

import (
	"ecommerce/config"
	"ecommerce/pkg/infrastructure/security"
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	config.LoadConfig()
	token, err := security.GenerateAccessToken("123")
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
func TestVerifyJwt(t *testing.T) {
	config.LoadConfig()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMyIsImV4cCI6MTczNzA0Mjk1OH0.l-YsGdRZkq1lZnxPZs3zQzg2YkTbKlvQ-Kym4_HPE8o"
	userId, err := security.VerifyJwt(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(userId)
}
