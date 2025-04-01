package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/api/v1/middleware"
	"server/internal/api/v1/websocket"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InviteService interface {
	InsertInvite(*entity.Invite, string) error
	UpdateStatusInvite(*entity.Invite, string, string) error
}

type inviteService struct {
	inviteRepository       repository.InviteRepository
	contactRepository      repository.ContactRepository
	notificationRepository repository.NotificationRepository
}

func NewInviteService(inviteRepository repository.InviteRepository, contactRepository repository.ContactRepository, notificationRepository repository.NotificationRepository) InviteService {
	return &inviteService{inviteRepository, contactRepository, notificationRepository}
}

const internalServerError = "error internal server"

func (service *inviteService) InsertInvite(invite *entity.Invite, cookieToken string) error {
	logger.Info("Validando o registro do convite")
	if err := invite.ValidateRegisterInvite(); err != nil {
		logger.Error("Erro na validação do convite: %v", err)
		return err
	}
	data, err := service.decodeToken(cookieToken)
	if err != nil {
		return err
	}
	id, err := service.extractUserID(data)
	if err != nil {
		return err
	}
	invite.UserIdInviter = id
	err = service.setInviteCreatedAt(invite)
	if err != nil {
		return err
	}
	if err := service.inviteRepository.InsertInvite(invite); err != nil {
		logger.Error("Erro ao inserir o convite no banco de dados: %v", err)
		return fmt.Errorf(internalServerError)
	}
	logger.Info("Sucesso ao inserir o convite no banco de dados!")
	username, err := service.extractUsername(data)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("<b>%s</b> enviou uma solicitação de contato.", username)
	err = service.insertNotification(invite, message)
	if err != nil {
		return err
	}
	err = websocket.SendMessageToUser(invite.UserIdInvited.Hex(), message, "notification")
	if err != nil {
		logger.Error("Erro ao enviar a notificação via WebSocket: %v", err)
	}
	return nil
}

func (service *inviteService) UpdateStatusInvite(invite *entity.Invite, statusInvite string, cookieToken string) error {
	logger.Info("Atualizando o status do convite")
	data, err := service.decodeToken(cookieToken)
	if err != nil {
		return err
	}
	id, err := service.extractUserID(data)
	if err != nil {
		return err
	}
	logger.Info("Verificando se existe convites entre os usuários")
	var userIDs = []primitive.ObjectID{invite.UserIdInvited, invite.UserIdInviter}
	invites, err := service.inviteRepository.FindInvitesByUsers(id, userIDs, "")
	if err != nil {
		logger.Error("Err ao buscar os convites: %v", err)
		return err
	}
	inviteData := invites[0]
	inviteID, err := service.convertStringToObjectID(inviteData.ID)
	if err != nil {
		return err
	}
	if statusInvite == "none" {
		logger.Info("Deletando o convite...")
		err = service.inviteRepository.DeleteInviteById(inviteID)
	} else if statusInvite == "added" {
		logger.Info("Adicionando contato e deletando o convite...")
		err = service.insertContact(invite)
		if err != nil {
			return err
		}
		err = service.inviteRepository.DeleteInviteById(inviteID)
	} else {
		logger.Info("Atualizando o status do convite...")
		err = service.inviteRepository.UpdateStatusInvite(inviteID, statusInvite)
	}
	if err != nil {
		return fmt.Errorf(internalServerError)
	}
	logger.Info("Sucesso ao atualizar o convite!")
	return nil
}

func (service *inviteService) insertContact(invite *entity.Invite) error {
	timestamp := time.Now().UnixMilli()
	contact := &entity.Contact{
		Status:       invite.Status,
		UserIdTarget: invite.UserIdInvited,
		UserIdActor:  invite.UserIdInviter,
		CreatedAt:    timestamp,
		UpdatedAt:    timestamp,
	}
	err := service.contactRepository.InsertContact(contact)
	if err != nil {
		logger.Error("Erro ao adicionar o contato: %v", err)
		return fmt.Errorf(internalServerError)
	}
	logger.Info("Contato adicionado com sucesso!")
	return nil
}

func (service *inviteService) insertNotification(invite *entity.Invite, message string) error {
	logger.Info("Registrando a notificação...")
	timestamp := time.Now().UnixMilli()
	notification := &entity.Notification{
		Message:      message,
		UserIdTarget: invite.UserIdInvited,
		UserIdActor:  invite.UserIdInviter,
		Type:         "invite",
		CreatedAt:    timestamp,
	}
	err := service.notificationRepository.InsertNotification(notification)
	if err != nil {
		logger.Error("Erro ao registrar a notificação: %v", err)
		return err
	}
	logger.Info("Notificação registrada com sucesso!")
	return nil
}

func (service *inviteService) decodeToken(cookieToken string) (primitive.M, error) {
	logger.Info("Obtendo informações armazenadas no cookie")
	data, err := middleware.NewMiddlewareToken().DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o cookie: %v", err)
		return primitive.M{}, fmt.Errorf(internalServerError)
	}
	return data, nil
}

func (service *inviteService) extractUsername(data primitive.M) (string, error) {
	logger.Info("Extraindo username do usuário que enviou o convite")
	username, ok := data["username"].(string)
	if !ok {
		logger.Error("Nome de usuário ausente ou inválido no token")
		return "", fmt.Errorf(internalServerError)
	}
	logger.Info("Nome de usuário extraído com sucesso.")
	return username, nil
}

func (service *inviteService) extractUserID(data primitive.M) (primitive.ObjectID, error) {
	logger.Info("Extraindo ID do usuário que enviou o convite")
	id, ok := data["id"].(string)
	if !ok {
		logger.Error("ID do usuário ausente ou inválido no token")
		return primitive.ObjectID{}, fmt.Errorf(internalServerError)
	}
	objectID, err := service.convertStringToObjectID(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objectID, nil
}

func (service *inviteService) convertStringToObjectID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro ao converter ID (string) para ObjectID: %v", err)
		return primitive.ObjectID{}, fmt.Errorf(internalServerError)
	}
	return objectID, nil
}

func (service *inviteService) setInviteCreatedAt(invite *entity.Invite) error {
	date, err := time.Parse(time.RFC3339, invite.CreatedAt)
	if err != nil {
		logger.Error("Erro ao fazer o parse da data: %v", err)
		return err
	}
	logger.Info("Inserindo o convite no banco de dados...")
	invite.CreatedAtAtMilliseconds = date.UnixMilli()
	invite.CreatedAt = ""
	return nil
}
