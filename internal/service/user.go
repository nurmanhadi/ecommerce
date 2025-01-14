package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/pkg/infrastructure/dto"
	"ecommerce/pkg/infrastructure/exception"
	"ecommerce/pkg/infrastructure/security"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type userImpl struct {
	userRepository repository.UserRepository
	validator      *validator.Validate
}

func NewUserService(userRepository *repository.UserRepository, validator *validator.Validate) UserService {
	return &userImpl{userRepository: *userRepository, validator: validator}
}
func (s *userImpl) UserRegister(req *dto.UserRegisterRequestDto) error {
	if err := s.validator.Struct(req); err != nil {
		return &exception.ValidationError{Message: err.Error()}
	}
	count, err := s.userRepository.CountUser(&req.Email)
	if err != nil {
		return err
	}
	if count == 1 {
		return exception.UserEmailAlreadyExist
	}
	hashPassword, err := security.HashPasswordBcrypt(req.Password)
	if err != nil {
		return err
	}

	id := uuid.NewString()
	user := &model.User{
		Id:        id,
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.userRepository.RegisterUser(user); err != nil {
		return err
	}
	return nil
}
func (s *userImpl) UserLogin(req *dto.UserLoginRequestDto) (map[string]interface{}, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}
	user, err := s.userRepository.GetUserByEmail(&req.Email)
	if err != nil {
		return nil, exception.UserEmailAndPasswordIsWrong
	}
	err = security.CompareFromPassword(user.Password, req.Password)
	if err != nil {
		return nil, exception.UserEmailAndPasswordIsWrong
	}
	token, err := security.GenerateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"access_token": token}, nil
}
