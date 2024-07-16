package controller

import (
	"net/http"
	"server/src/config"
	"server/src/model/dto"
	"server/src/model/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(user service.UserService) UserController {
	return &userController{user}
}

var logger *config.Logger = config.NewLogger("controller")

func (controller *userController) CreateUser(ctx *gin.Context) {
	method := ctx.Request.Method
	url := ctx.Request.URL
	remoteAddr := ctx.Request.RemoteAddr
	logger.Info("(%s - %s) %s Criando usuário...", method, url, remoteAddr)
	var user *dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error("Erro ao converter o JSON: %v", err)
		panic(err)
	}

	logger.Info("Enviando o usuário para validação...")
	if err := controller.userService.CreateUser(user); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusCreated, "Usuário criado com sucesso!")
}
