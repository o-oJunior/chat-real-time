package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/api/v1/middleware"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactService interface {
	UpdateStatusContact(*entity.Contact, string, string) error
}

type contactService struct {
	contactRepository repository.ContactRepository
}

func NewContactService(contactRepository repository.ContactRepository) ContactService {
	return &contactService{contactRepository}
}

func (service *contactService) UpdateStatusContact(contact *entity.Contact, statusContact string, cookieToken string) error {
	logger.Info("Atualizando o status do contato")
	middleware := middleware.NewMiddlewareToken()
	data, err := middleware.DecodeToken(cookieToken)
	if err != nil {
		logger.Error("Erro ao decodificar o cookie: %v", err)
		return err
	}
	idString := data["id"].(string)
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		logger.Error("Erro ao converter o ID (string) para ObjectID")
		return fmt.Errorf("error internal server")
	}
	logger.Info("Verificando se existe relação de contato entre os usuários")
	var userIDs = []primitive.ObjectID{contact.UserIdTarget, contact.UserIdActor}
	contacts, err := service.contactRepository.FindContactByUsers(id, userIDs)
	if err != nil {
		logger.Error("Erro ao buscar o contato: %v", err)
		return err
	}
	if len(contacts) == 0 {
		logger.Warn("Nenhum contato encontrado!")
		return fmt.Errorf("contact not found")
	}
	contactData := contacts[0]
	if statusContact == "none" {
		err = service.contactRepository.DeleteContactById(contactData.ID)
		if err != nil {
			logger.Error("Erro ao deletar o contato: %v", err)
			return fmt.Errorf("error internal server")
		}
		logger.Info("Sucesso ao excluir o contato!")
	} else if statusContact == "blocked" {
		timestamp := time.Now().UnixMilli()
		err = service.contactRepository.UpdateStatusContact(id, statusContact, timestamp)
		if err != nil {
			logger.Error("Erro ao atualizar o contato: %v", err)
			return fmt.Errorf("error internal server")
		}
		logger.Info("Sucesso ao atualizar o contato!")
	} else {
		return fmt.Errorf("contact status invalid")
	}
	return nil
}
