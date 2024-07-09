package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"server/src/config/logger"
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

func (uc userController) CreateUser(context *gin.Context) {
	method := context.Request.Method
	url := context.Request.URL
	remoteAddr := context.Request.RemoteAddr
	logger.Info("[CONTROLLER (%s - %s)] %s Criando usuário...", method, url, remoteAddr)
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		logger.Error("Erro ao ler o dados enviados!!", err, false)
	}
	var user dto.UserDTO
	if err = json.Unmarshal(body, &user); err != nil {
		logger.Error("Erro ao converter o JSON", err, false)
	}
	logger.Info("[CONTROLLER] Enviando o usuário para validação...")
	err = uc.userService.CreateUser(user)
	if err != nil {
		logger.Error("[CONTROLLER] Erro ao criar o usuário!", err, false)
		context.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"message":    "Ocorreu um erro ao criar o usuário!",
		})
		return
	}
	logger.Info("[CONTROLLER] Usuário validado e criado com sucesso!")
	context.JSON(http.StatusCreated, gin.H{
		"statusCode": 201,
		"message":    "Usuário criado com sucesso!",
	})
}
