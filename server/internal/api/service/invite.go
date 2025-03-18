package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/api/v1/middleware"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InviteService interface {
	InsertInvite(*entity.Invite, string) error
	UpdateStatusInvite(*entity.Invite, string, string) error
}

type inviteService struct {
	inviteRepository  repository.InviteRepository
	contactRepository repository.ContactRepository
}

func NewInviteService(inviteRepository repository.InviteRepository, contactRepository repository.ContactRepository) InviteService {
	return &inviteService{inviteRepository, contactRepository}
}

func (service *inviteService) InsertInvite(invite *entity.Invite, cookieToken string) error {
	logger.Info("Validando o registro do convite")
	if err := invite.ValidateRegisterInvite(); err != nil {
		logger.Error("Erro na validação do convite: %v", err)
		return err
	}
	logger.Info("Obtendo informações armazenadas no cookie")
	middleware := middleware.NewMiddlewareToken()
	data, err := middleware.DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o cookie: %v", err)
		return err
	}
	logger.Info("Extraindo ID do usuário que enviou o convite")
	id, ok := data["id"].(string)
	if !ok {
		logger.Error("ID do usuário ausente ou inválido no token")
		return fmt.Errorf("error internal server")
	}
	objectID, err := service.convertStringToObjectID(id)
	if err != nil {
		return err
	}
	invite.UserIdInviter = objectID
	date, err := time.Parse(time.RFC3339, invite.CreatedAt)
	if err != nil {
		logger.Error("Erro ao fazer o parse da data: %v", err)
		return err
	}
	logger.Info("Inserindo o convite no banco de dados...")
	invite.CreatedAtAtMilliseconds = date.UnixMilli()
	invite.CreatedAt = ""
	if err := service.inviteRepository.InsertInvite(invite); err != nil {
		logger.Error("Erro ao inserir o convite no banco de dados: %v", err)
		return fmt.Errorf("error internal server")
	}
	logger.Info("Sucesso ao inserir o convite no banco de dados!")
	return nil
}

func (service *inviteService) UpdateStatusInvite(invite *entity.Invite, statusInvite string, cookieToken string) error {
	logger.Info("Atualizando o status do convite")
	middleware := middleware.NewMiddlewareToken()
	data, err := middleware.DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o cookie: %v", err)
		return err
	}
	idString := data["id"].(string)
	logger.Info("Verificando se existe convites entre os usuários")
	var userIDs = []primitive.ObjectID{invite.UserIdInvited, invite.UserIdInviter}
	objectID, err := service.convertStringToObjectID(idString)
	if err != nil {
		return err
	}
	invites, err := service.inviteRepository.FindInvitesByUsers(objectID, userIDs, "")
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
			return fmt.Errorf("error internal server")
		}
		err = service.inviteRepository.DeleteInviteById(inviteID)
	} else {
		logger.Info("Atualizando o status do convite...")
		err = service.inviteRepository.UpdateStatusInvite(inviteID, statusInvite)
	}
	if err != nil {
		return fmt.Errorf("error internal server")
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
		return err
	}
	logger.Info("Contato adicionado com sucesso!")
	return nil
}

func (service *inviteService) convertStringToObjectID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro ao converter ID (string) para ObjectID: %v", err)
		return primitive.ObjectID{}, fmt.Errorf("error internal server")
	}
	return objectID, nil
}
