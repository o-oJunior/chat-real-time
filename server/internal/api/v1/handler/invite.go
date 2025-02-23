package handler

import (
	"net/http"
	"server/internal/api/entity"
	"server/internal/api/service"
	"server/internal/api/v1/response"

	"github.com/gin-gonic/gin"
)

type InviteHandler interface {
	InsertInvite(*gin.Context)
	UpdateStatusInvite(*gin.Context)
}

type inviteHandler struct {
	inviteService service.InviteService
}

func NewInviteHandler(service service.InviteService) InviteHandler {
	return &inviteHandler{service}
}

func (handler *inviteHandler) converterJsonInvite(ctx *gin.Context, message string) *entity.Invite {
	method := ctx.Request.Method
	url := ctx.Request.URL
	remoteAddr := ctx.Request.RemoteAddr
	logger.Info("(%s - %s) %s %s", method, url, remoteAddr, message)
	var invite *entity.Invite
	if err := ctx.ShouldBindJSON(&invite); err != nil {
		logger.Error("Erro ao converter o JSON: %v", err)
		panic(err)
	}
	return invite
}

func (handler *inviteHandler) InsertInvite(ctx *gin.Context) {
	cookieToken, err := ctx.Cookie("token")
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, "access unauthorized")
		return
	}
	invite := handler.converterJsonInvite(ctx, "Registrando o convite...")
	err = handler.inviteService.InsertInvite(invite, cookieToken)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SendSuccess(ctx, http.StatusOK, "invitation sent successfully", nil)
}

func (handler *inviteHandler) UpdateStatusInvite(ctx *gin.Context) {
	cookieToken, err := ctx.Cookie("token")
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, "access unauthorized")
		return
	}
	invite := handler.converterJsonInvite(ctx, "Obtendo as informações do usuário...")
	statusInvite := ctx.Param("status")
	err = handler.inviteService.UpdateStatusInvite(invite, statusInvite, cookieToken)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SendSuccess(ctx, http.StatusOK, "invitation successfully updated", nil)
}
