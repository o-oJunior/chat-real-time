package handler

import (
	"fmt"
	"net/http"
	"server/internal/api/entity"
	"server/internal/api/service"
	"server/internal/api/v1/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactHandler interface {
	UpdateStatusContact(*gin.Context)
}

type contactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(contactService service.ContactService) ContactHandler {
	return &contactHandler{contactService}
}

func (handler *contactHandler) converterJsonContact(ctx *gin.Context, message string) (*entity.Contact, error) {
	method := ctx.Request.Method
	url := ctx.Request.URL
	remoteAddr := ctx.Request.RemoteAddr
	logger.Info("(%s - %s) %s %s", method, url, remoteAddr, message)
	var invite *entity.Invite
	if err := ctx.ShouldBindJSON(&invite); err != nil {
		logger.Error("Erro ao converter o JSON: %v", err)
		panic(err)
	}
	userIdInviter, err := primitive.ObjectIDFromHex(invite.UserIdInviter)
	if err != nil {
		logger.Error("Erro ao converter ID do usuário que enviou o convite: %v", err)
		return nil, fmt.Errorf("error internal server")
	}
	userIdInvited, err := primitive.ObjectIDFromHex(invite.UserIdInvited)
	if err != nil {
		logger.Error("Erro ao converter ID do usuário convidado: %v", err)
		return nil, fmt.Errorf("error internal server")
	}
	contact := &entity.Contact{
		Status:       invite.InviteStatus,
		UserIdTarget: userIdInvited,
		UserIdActor:  userIdInviter,
	}
	return contact, nil
}

func (handler *contactHandler) UpdateStatusContact(ctx *gin.Context) {
	cookieToken, err := ctx.Cookie("token")
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, "access unauthorized")
		return
	}
	contact, err := handler.converterJsonContact(ctx, "Obtendo as informações do contato...")
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	statusContact := ctx.Param("status")
	err = handler.contactService.UpdateStatusContact(contact, statusContact, cookieToken)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SendSuccess(ctx, http.StatusOK, "contact successfully updated", nil)
}
