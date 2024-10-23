package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/api/v1/middleware"
	"server/internal/config"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService interface {
	GetUsersExceptID(string, string, *options.FindOptions) (*[]entity.User, int64, error)
	CreateUser(*entity.User) error
	Authentication(*entity.User) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user}
}

var logger *config.Logger = config.NewLogger("service")

func (service *userService) GetUsersExceptID(username string, cookieToken string, options *options.FindOptions) (*[]entity.User, int64, error) {
	middlewareToken := middleware.NewMiddlewareToken()
	data, err := middlewareToken.DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o token: %v", err)
		return nil, 0, err
	}
	idString := data["id"].(string)
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		logger.Error("Erro ao converter o ID string para ObjectID: %v", err)
		return nil, 0, err
	}
	users, totalUsers, err := service.userRepository.GetUsersAndTotalExceptID(id, username, options)
	if err != nil {
		logger.Error("Erro ao buscar os usuários: %v", err)
		return nil, 0, err
	}
	return users, totalUsers, nil
}

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
		err := fmt.Errorf("nome de usuário já cadastrado")
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

func (userService *userService) Authentication(user *entity.User) (*entity.User, error) {
	logger.Info("Buscando o usuário...")
	data, err := userService.userRepository.FindUsername(user)
	logger.Info("Validando a credenciais...")
	if err != nil {
		logger.Error("Erro ao autenticar o usuário: %v", err)
		return nil, fmt.Errorf("usuário e/ou senha inválido(s)")
	}
	if err := user.ComparePassword(data.HashPassword); err != nil {
		logger.Error("Erro ao validar a senha: %v", err)
		return nil, fmt.Errorf("usuário e/ou senha inválido(s)")
	}
	logger.Info("Credenciais válidas")
	data.HashPassword = ""
	data.CreatedAt = time.UnixMilli(data.CreatedAtMilliseconds).UTC().Format(time.RFC3339)
	return data, nil
}
