package service

import (
	"server/src/config/logger"
	"server/src/model/dto"
	"server/src/model/repository"
)

type UserService interface {
	CreateUser(user dto.UserDTO) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user}
}

func (us *userService) CreateUser(user dto.UserDTO) error {
	logger.Info("[SERVICE] Validando usuário...")
	err := us.userRepository.CreateUser(user)
	return err
}
