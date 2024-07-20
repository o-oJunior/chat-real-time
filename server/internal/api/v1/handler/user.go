package handler

import (
	"net/http"
	"server/internal/api/entity"
	"server/internal/api/service"
	"server/internal/config"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(*gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(user service.UserService) UserHandler {
	return &userHandler{user}
}

var logger *config.Logger = config.NewLogger("controller")

func (controller *userHandler) CreateUser(ctx *gin.Context) {
	method := ctx.Request.Method
	url := ctx.Request.URL
	remoteAddr := ctx.Request.RemoteAddr
	logger.Info("(%s - %s) %s Criando usuário...", method, url, remoteAddr)
	var user *entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error("Erro ao converter o JSON: %v", err)
		panic(err)
	}

	logger.Info("Enviando o usuário para validação...")
	if err := controller.userService.CreateUser(user); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusCreated, "Usuário criado com sucesso!", nil)
}
