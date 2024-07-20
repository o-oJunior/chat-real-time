package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/config"
	"strings"
)

type UserService interface {
	CreateUser(user *entity.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user}
}

var logger *config.Logger = config.NewLogger("service")

func (userService *userService) CreateUser(user *entity.User) error {
	logger.Info("Validando usuário...")
	if err := user.ValidateCreateUser(); err != nil {
		logger.Error("Erro ao validar o usuário: %v", err)
		return err
	}
	data, _ := userService.userRepository.FindUsername(user)
	dataUserNameLower := strings.ToLower(data.Username)
	userNameLower := strings.ToLower(user.Username)
	if dataUserNameLower == userNameLower {
		err := fmt.Errorf("username já está cadastrado")
		logger.Error("Erro ao cadastrar o usuário: %v", err)
		return err
	}
	logger.Info("O usuário não existe, a conta será criada")
	if err := user.EncodePassword(); err != nil {
		logger.Error("Erro ao criptografar a senha %v", err)
		return err
	}
	err := userService.userRepository.InsertUser(user)
	return err
}
