package service

import (
	"fmt"
	"server/src/config"
	"server/src/model/dto"
	"server/src/model/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *dto.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user}
}

var logger *config.Logger = config.NewLogger("service")

func (userService *userService) CreateUser(user *dto.User) error {
	logger.Info("Validando usuário...")
	if err := userService.validateCreateUser(user); err != nil {
		logger.Error("Erro ao validar o usuário -> %v", err)
		return err
	}
	logger.Info("Usuário válido")
	logger.Info("Em processo de criptografia de senha...")
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		logger.Error("Erro na tentativa de criptografar a senha: %v", err)
		return err
	}
	logger.Info("Senha criptografada!")
	user.HashPassword = string(hashPassword)
	user.Password = ""
	err = userService.userRepository.CreateUser(user)
	return err
}

func errorParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (service *userService) validateCreateUser(user *dto.User) error {
	if user.Username == "" && user.FirstName == "" &&
		user.LastName == "" && user.Email == "" && user.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if user.Username == "" {
		return errorParamIsRequired("username", "string")
	}

	if user.FirstName == "" {
		return errorParamIsRequired("firstName", "string")
	}

	if user.LastName == "" {
		return errorParamIsRequired("lastName", "string")
	}

	if user.Email == "" {
		return errorParamIsRequired("email", "string")
	}

	if user.Password == "" {
		return errorParamIsRequired("password", "string")
	}

	return nil
}
