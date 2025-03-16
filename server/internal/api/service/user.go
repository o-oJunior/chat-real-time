package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/api/v1/middleware"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUsersExceptID(string, string, *middleware.Pagination) (*[]entity.User, int, error)
	CreateUser(*entity.User) error
	Authentication(*entity.User) (*entity.User, error)
	GetContacts(string, *middleware.Pagination, string, string) (*[]entity.User, int, error)
}

type userService struct {
	userRepository    repository.UserRepository
	inviteRepository  repository.InviteRepository
	contactRepository repository.ContactRepository
}

func NewUserService(user repository.UserRepository, invite repository.InviteRepository, contact repository.ContactRepository) UserService {
	return &userService{userRepository: user, inviteRepository: invite, contactRepository: contact}
}

func (service *userService) GetUsersExceptID(username, cookieToken string, pagination *middleware.Pagination) (*[]entity.User, int, error) {
	logger.Info("Decodificando token...")
	data, err := middleware.NewMiddlewareToken().DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o token: %v", err)
		return nil, 0, fmt.Errorf("access unauthorized")
	}
	idString, ok := data["id"].(string)
	if !ok {
		logger.Error("ID do usuário ausente ou inválido no token")
		return nil, 0, fmt.Errorf("error internal server")
	}
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		logger.Error("Erro ao converter o ID string para ObjectID: %v", err)
		return nil, 0, err
	}
	if username == "" {
		logger.Warn("Busca de usuários sem filtro de 'username', irá trazer %d usuários", pagination.Limit)
	}
	logger.Info("Aplicando filtros e buscando usuários...")
	filter := bson.M{
		"_id": bson.M{"$ne": id},
		"username": bson.M{
			"$regex": primitive.Regex{Pattern: "^" + username, Options: "i"},
		},
	}
	users, totalUsers, err := service.userRepository.GetUsersWithFilter(filter, pagination)
	if err != nil {
		logger.Error("Erro ao buscar os usuários: %v", err)
		return nil, 0, err
	}
	return service.mapInvitesAndContactsToUsers(idString, users, totalUsers)
}

func (service *userService) mapInvitesAndContactsToUsers(userID string, users *[]entity.User, totalUsers int) (*[]entity.User, int, error) {
	logger.Info("Verificando convites e contatos entre os usuários...")
	userIDs := make([]string, len(*users))
	for i, user := range *users {
		userIDs[i] = user.ID
	}
	logger.Info("Buscando convites com base nos IDs dos usuários")
	invites, err := service.inviteRepository.FindInvitesByUsers(userID, userIDs, "")
	if err != nil {
		logger.Error("Erro ao buscar os convites com base nos IDs dos usuários: %v", err)
		return nil, 0, err
	}
	inviteMap := service.mountInviteMap(userID, invites)
	contacts, err := service.contactRepository.GetContactsByUser(userID)
	if err != nil {
		logger.Error("Erro ao buscar os contatos com base no ID do usuário logado: %v", err)
		return nil, 0, err
	}
	contactMap := service.mountContactMap(userID, contacts)
	for i := range *users {
		user := &(*users)[i]
		if invite, exists := inviteMap[user.ID]; exists {
			user.InviteStatus = invite.InviteStatus
			user.UserIdInvited = invite.UserIdInvited
			user.UserIdInviter = invite.UserIdInviter
		} else if contact, exists := contactMap[user.ID]; exists {
			user.InviteStatus = contact.InviteStatus
		} else {
			user.InviteStatus = ""
		}
	}
	logger.Info("Retornando %d usuários...", len(*users))
	return users, totalUsers, nil
}

func (service *userService) mapInvitesToUsers(userID string, users *[]entity.User, totalUsers int) (*[]entity.User, int, error) {
	logger.Info("Verificando convites entre os usuários...")
	userIDs := make([]string, len(*users))
	for i, user := range *users {
		userIDs[i] = user.ID
	}
	logger.Info("Buscando convites com base nos IDs dos usuários")
	invites, err := service.inviteRepository.FindInvitesByUsers(userID, userIDs, "")
	if err != nil {
		logger.Error("Erro ao buscar os convites com base nos IDs dos usuários")
		return nil, 0, err
	}
	inviteMap := service.mountInviteMap(userID, invites)
	for i := range *users {
		user := &(*users)[i]
		if invite, exists := inviteMap[user.ID]; exists {
			user.InviteStatus = invite.InviteStatus
			user.UserIdInvited = invite.UserIdInvited
			user.UserIdInviter = invite.UserIdInviter
		} else {
			user.InviteStatus = ""
		}
	}
	logger.Info("Retornando %d usuários...", len(*users))
	return users, totalUsers, nil
}

func (service *userService) mountInviteMap(userID string, invites []entity.Invite) map[string]entity.Invite {
	logger.Info("Iterando convites para os usuários")
	inviteMap := make(map[string]entity.Invite)
	for _, invite := range invites {
		targetID := invite.UserIdInvited
		if invite.UserIdInviter != userID {
			targetID = invite.UserIdInviter
		}
		inviteMap[targetID] = invite
	}
	return inviteMap
}

func (service *userService) mountContactMap(userID string, contacts []entity.Contact) map[string]entity.Contact {
	logger.Info("Iterando contatos para os usuários")
	contactMap := make(map[string]entity.Contact)
	for _, contact := range contacts {
		targetID := contact.UserIdInvited.Hex()
		if contact.UserIdInviter.Hex() != userID {
			targetID = contact.UserIdInviter.Hex()
		}
		contactMap[targetID] = contact
	}
	return contactMap
}

func (userService *userService) CreateUser(user *entity.User) error {
	logger.Info("Validando e criando usuário...")
	if err := user.ValidateCreateUser(); err != nil {
		logger.Error("Erro ao validar o usuário: %v", err)
		return err
	}
	logger.Info("Procurando se o usuário já está cadastrado no banco")
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
	if err != nil {
		logger.Error("Erro ao inserir o usuário no banco de dados: %v", err)
		return err
	}
	return nil
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

func (service *userService) GetContacts(cookieToken string, pagination *middleware.Pagination, group, username string) (*[]entity.User, int, error) {
	logger.Info("Obtendo informações armazenadas no cookie")
	data, err := middleware.NewMiddlewareToken().DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o cookie: %v", err)
		return nil, 0, fmt.Errorf("access unauthorized")
	}
	userIdLogged, ok := data["id"].(string)
	if !ok {
		logger.Error("ID do usuário ausente ou inválido no token")
		return nil, 0, fmt.Errorf("error internal server")
	}
	if group == "added" {
		return service.getAddedContacts(userIdLogged, username, pagination)
	}
	searchField, validGroup := map[string]string{
		"received": "userIdInvited",
		"sent":     "userIdInviter",
	}[group]
	if !validGroup {
		logger.Error("Grupo inválido: %s", group)
		return nil, 0, fmt.Errorf("invalid group")
	}
	userIDs, err := service.getUserIDsFromInvites(userIdLogged, searchField)
	if err != nil {
		logger.Error("Error ao buscar os usuários a partir dos convites: %v", err)
		return nil, 0, err
	}
	if len(userIDs) == 0 {
		logger.Warn("Nenhum usuário a partir dos convites")
		return nil, 0, nil
	}
	filter := service.mountFilterByUserIDs(userIDs, username)
	users, totalUsers, err := service.userRepository.GetUsersWithFilter(filter, pagination)
	if err != nil {
		logger.Error("Erro ao buscar os usuários: %v", err)
		return nil, 0, err
	}
	return service.mapInvitesToUsers(userIdLogged, users, totalUsers)
}

func (service *userService) getAddedContacts(userIdLogged, username string, pagination *middleware.Pagination) (*[]entity.User, int, error) {
	logger.Info("Buscando os contatos do usuário logado")
	contacts, err := service.contactRepository.GetContactsByUser(userIdLogged)
	if err != nil {
		logger.Error("Erro ao obter o usuário através do ID: %v", err)
		return nil, 0, err
	}
	if len(contacts) == 0 {
		logger.Warn("O usuário não tem nenhum contato!")
		return nil, 0, nil
	}
	objectIDLogged, err := primitive.ObjectIDFromHex(userIdLogged)
	if err != nil {
		logger.Error("Erro ao converter ID do usuário logado: %v", err)
		return nil, 0, err
	}
	logger.Info("Acessando os convites para extrair os IDs dos usuário")
	var userIDs []primitive.ObjectID
	for _, contact := range contacts {
		userID := contact.UserIdInvited
		if objectIDLogged != contact.UserIdInviter {
			userID = contact.UserIdInviter
		}
		userIDs = append(userIDs, userID)
	}
	logger.Info("Retornando %d IDs", len(userIDs))
	filter := service.mountFilterByUserIDs(userIDs, username)
	logger.Info("Obtendo as informações de %d contatos", len(contacts))
	users, totalUsers, err := service.userRepository.GetUsersWithFilter(filter, pagination)
	if err != nil {
		logger.Error("Erro ao buscar os usuários adicionados como contato: %v", err)
		return nil, 0, err
	}
	for i := range *users {
		(*users)[i].InviteStatus = contacts[i].InviteStatus
	}
	return users, totalUsers, nil
}

func (service *userService) getUserIDsFromInvites(userIdLogged, searchField string) ([]primitive.ObjectID, error) {
	logger.Info("Filtrando convites pelo campo '%s'", searchField)
	invites, err := service.inviteRepository.FindInvitesByUsers(userIdLogged, []string{""}, searchField)
	if err != nil {
		return nil, err
	}
	logger.Info("Acessando os convites para extrair os IDs dos usuário")
	var userIDs []primitive.ObjectID
	for _, invite := range invites {
		userID := invite.UserIdInvited
		if searchField == "userIdInvited" {
			userID = invite.UserIdInviter
		}
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			logger.Error("Erro ao converter ID do usuário: %v", err)
			return nil, err
		}
		userIDs = append(userIDs, objectID)
	}
	logger.Info("Retornando %d IDs", len(userIDs))
	return userIDs, nil
}

func (service *userService) mountFilterByUserIDs(userIDs []primitive.ObjectID, username string) bson.M {
	return bson.M{
		"_id": bson.M{"$in": userIDs},
		"username": bson.M{
			"$regex": primitive.Regex{Pattern: "^" + username, Options: "i"},
		},
	}
}
