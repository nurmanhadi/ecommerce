package service

import "ecommerce/pkg/infrastructure/dto"

type UserService interface {
	UserRegister(req *dto.UserRegisterRequestDto) error
	UserLogin(req *dto.UserLoginRequestDto) (map[string]interface{}, error)
}
