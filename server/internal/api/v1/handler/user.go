package handler

import (
	"net/http"
	"server/internal/api/entity"
	"server/internal/api/service"
	"server/internal/api/v1/middleware"
	"server/internal/api/v1/response"
	"server/internal/config"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(*gin.Context)
	Authentication(*gin.Context)
	Logout(*gin.Context)
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
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SendSuccess(ctx, http.StatusCreated, "Usuário criado com sucesso!", nil)
}

func (handler *userHandler) Authentication(ctx *gin.Context) {
	user := handler.converterJSON(ctx, "Autenticando usuário...")
	logger.Info("Enviando o usuário para a validação...")
	data, err := handler.userService.Authentication(user)
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	newMiddlewareToken := middleware.NewMiddlewareToken()
	token, err := newMiddlewareToken.Generate(data)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, err.Error())
	}
	logger.Info("Token gerado, será armazenado nos cookies")
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("token", token, int(time.Hour*24), "/", "", true, true)
	response.SendSuccess(ctx, http.StatusOK, "login efetuado com sucesso!", nil)
}

func (handler *userHandler) Logout(ctx *gin.Context) {
	logger.Info("Realizando o logout")
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("token", "", -1, "/", "", true, true)
	response.SendSuccess(ctx, http.StatusOK, "logout efetuado com sucesso!", nil)
}
