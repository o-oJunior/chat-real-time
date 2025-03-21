package handler

import (
	"math"
	"net/http"
	"server/internal/api/entity"
	"server/internal/api/service"
	"server/internal/api/v1/middleware"
	"server/internal/api/v1/response"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UserHandler interface {
	GetUsers(*gin.Context)
	GetUserToken(*gin.Context)
	CreateUser(*gin.Context)
	Authentication(*gin.Context)
	Logout(*gin.Context)
	GetContacts(*gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service}
}

func (handler *userHandler) converterJSON(ctx *gin.Context, message string) *entity.User {
	method, url, remoteAddr := ctx.Request.Method, ctx.Request.URL, ctx.Request.RemoteAddr
	logger.Info("(%s - %s) %s %s", method, url, remoteAddr, message)
	var user *entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error("Erro ao converter o JSON: %v", err)
		panic(err)
	}
	return user
}

func (handler *userHandler) handleError(ctx *gin.Context, err error, statusCode int) {
	if err.Error() == "access unauthorized" {
		response.SendError(ctx, http.StatusUnauthorized, err.Error())
	} else {
		response.SendError(ctx, statusCode, "internal server error")
	}
}

func (handler *userHandler) calculateTotalPages(totalUsers int, limit int) int {
	return int(math.Ceil(float64(totalUsers) / float64(limit)))
}

func (handler *userHandler) sendUserListResponse(ctx *gin.Context, users *[]entity.User, totalUsers int, pagination *middleware.Pagination) {
	totalPages := handler.calculateTotalPages(totalUsers, pagination.Limit)
	if users == nil || len(*users) == 0 {
		response.SendSuccess(ctx, http.StatusOK, "", bson.M{"users": []entity.User{}})
		return
	}
	response.SendSuccess(ctx, http.StatusOK, "", bson.M{
		"page":       pagination.Page,
		"totalPages": totalPages,
		"users":      *users,
	})
}

func (handler *userHandler) GetUsers(ctx *gin.Context) {
	pagination := middleware.ParsePagination(ctx)
	username := ctx.Query("username")
	cookieToken, err := ctx.Cookie("token")
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, "access unauthorized")
		return
	}
	logger.Info("Consultando página %d com limite %d de usuários", pagination.Page, pagination.Limit)
	users, totalUsers, err := handler.userService.GetUsersExceptID(username, cookieToken, pagination)
	if err != nil {
		handler.handleError(ctx, err, http.StatusBadRequest)
		return
	}
	handler.sendUserListResponse(ctx, users, totalUsers, pagination)
}

func (handler *userHandler) GetUserToken(ctx *gin.Context) {
	cookieToken, _ := ctx.Cookie("token")
	data, _ := middleware.NewMiddlewareToken().DecodeToken(cookieToken)
	response.SendSuccess(ctx, http.StatusOK, "", data)
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
		return
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

func (handler *userHandler) GetContacts(ctx *gin.Context) {
	cookieToken, err := ctx.Cookie("token")
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, "access unauthorized")
		return
	}
	pagination := middleware.ParsePagination(ctx)
	group := ctx.Query("group")
	username := ctx.Query("username")
	users, totalUsers, err := handler.userService.GetContacts(cookieToken, pagination, group, username)
	if err != nil {
		handler.handleError(ctx, err, http.StatusBadRequest)
		return
	}
	handler.sendUserListResponse(ctx, users, totalUsers, pagination)
}
