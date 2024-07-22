package handler

import (
	"net/http"
	"server/internal/api/entity"
	"server/internal/api/service"
	"server/internal/config"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(*gin.Context)
	Authentication(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(user service.UserService) UserHandler {
	return &userHandler{user}
}

var logger *config.Logger = config.NewLogger("handler")

func (handler *userHandler) converterJSON(ctx *gin.Context, message string) *entity.User {
	method := ctx.Request.Method
	url := ctx.Request.URL
	remoteAddr := ctx.Request.RemoteAddr
	logger.Info("(%s - %s) %s %s", method, url, remoteAddr, message)
	var user *entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error("Erro ao converter o JSON: %v", err)
		panic(err)
	}
	return user
}

func (handler *userHandler) CreateUser(ctx *gin.Context) {
	user := handler.converterJSON(ctx, "Criando usuário...")
	logger.Info("Enviando o usuário para validação...")
	if err := handler.userService.CreateUser(user); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(ctx, http.StatusCreated, "Usuário criado com sucesso!", nil)
}

func (handler *userHandler) Authentication(ctx *gin.Context) {
	user := handler.converterJSON(ctx, "Autenticando usuário...")
	logger.Info("Enviando o usuário para a validação...")
	data, err := handler.userService.Authentication(user)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.SetCookie("token", data.Token, int(time.Hour*24), "/", "", true, true)
	data.Token = ""
	sendSuccess(ctx, http.StatusOK, "", data)
}
