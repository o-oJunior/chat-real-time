package service

import (
	"fmt"
	"server/internal/api/entity"
	"server/internal/api/repository"
	"server/internal/api/v1/middleware"
	"time"
)

type InviteService interface {
	InsertInvite(*entity.Invite, string) error
	UpdateStatusInvite(*entity.Invite, string, string) error
}

type inviteService struct {
	inviteRepository repository.InviteRepository
}

func NewInviteService(inviteRepository repository.InviteRepository) InviteService {
	return &inviteService{inviteRepository}
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
	invite.UserIdInviter = id
	date, err := time.Parse(time.RFC3339, invite.InvitedAt)
	if err != nil {
		logger.Error("Erro ao fazer o parse da data: %v", err)
		return err
	}
	logger.Info("Inserindo o convite no banco de dados...")
	invite.InvitedAtMilliseconds = date.UnixMilli()
	invite.InvitedAt = ""
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
	var userIDs = []string{invite.UserIdInvited, invite.UserIdInviter}
	invites, err := service.inviteRepository.FindInvitesByUsers(idString, userIDs, "")
	if err != nil {
		return err
	}
	inviteData := invites[0]
	if statusInvite == "none" {
		logger.Info("Deletando o convite...")
		err = service.inviteRepository.DeleteInviteById(inviteData.ID)
	} else {
		logger.Info("Atualizando o status do convite...")
		err = service.inviteRepository.UpdateStatusInvite(inviteData.ID, statusInvite)
	}
	if err != nil {
		return fmt.Errorf("error internal server")
	}
	logger.Info("Sucesso ao atualizar o convite!")
	return nil
}
