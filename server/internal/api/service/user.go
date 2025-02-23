package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/api/v1/middleware"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUsersExceptID(string, string, int, int) (*[]entity.User, int64, error)
	CreateUser(*entity.User) error
	Authentication(*entity.User) (*entity.User, error)
}

type userService struct {
	userRepository   repository.UserRepository
	inviteRepository repository.InviteRepository
}

func NewUserService(user repository.UserRepository, invite repository.InviteRepository) UserService {
	return &userService{userRepository: user, inviteRepository: invite}
}

func (service *userService) GetUsersExceptID(username string, cookieToken string, limit int, offset int) (*[]entity.User, int64, error) {
	middlewareToken := middleware.NewMiddlewareToken()
	data, err := middlewareToken.DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o token: %v", err)
		return nil, 0, fmt.Errorf("access unauthorized")
	}
	idString := data["id"].(string)
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		logger.Error("Erro ao converter o ID string para ObjectID: %v", err)
		return nil, 0, err
	}
	users, totalUsers, err := service.userRepository.GetUsersAndTotalExceptID(id, username, limit, offset)
	if err != nil {
		logger.Error("Erro ao buscar os usuários: %v", err)
		return nil, 0, err
	}
	logger.Info("Verificando se existe convites entre os usuários")
	userIDs := make([]string, len(*users))
	for i, user := range *users {
		userIDs[i] = user.ID
	}

	invites, err := service.inviteRepository.FindInvitesByUsers(idString, userIDs)
	if err != nil {
		return nil, 0, err
	}

	inviteMap := make(map[string]entity.Invite)
	for _, invite := range invites {
		var userID string
		if invite.UserIdInviter != idString {
			userID = invite.UserIdInviter
		} else {
			userID = invite.UserIdInvited
		}
		inviteMap[userID] = entity.Invite{
			InviteStatus:  invite.InviteStatus,
			UserIdInvited: invite.UserIdInvited,
			UserIdInviter: invite.UserIdInviter,
		}
	}

	for i := range *users {
		user := &(*users)[i]
		if value, exists := inviteMap[user.ID]; exists {
			user.InviteStatus = value.InviteStatus
			user.UserIdInvited = value.UserIdInvited
			user.UserIdInviter = value.UserIdInviter
		} else {
			user.InviteStatus = ""
		}
	}

	return users, totalUsers, nil
}

func (userService *userService) CreateUser(user *entity.User) error {
	logger.Info("Validando usuário...")
	if err := user.ValidateCreateUser(); err != nil {
		logger.Error("Erro ao validar o usuário: %v", err)
		return err
	}
	data, _ := userService.userRepository.FindUsername(user.Username)
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
	data, err := userService.userRepository.FindUsername(user.Username)
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
