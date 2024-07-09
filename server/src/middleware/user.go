package middleware

import (
	"server/src/config/logger"
	"server/src/config/mongodb"
	"server/src/controller"
	"server/src/model/repository"
	"server/src/model/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMiddleware interface {
	CreateUser(*gin.Context)
}

type userMiddleware struct {
	userController controller.UserController
}

func NewUserMiddleware() UserMiddleware {
	return &userMiddleware{}
}

func (userMiddleware) handleConnection() (controller.UserController, *mongo.Database) {
	database := mongodb.Connect()
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController, database
}

func (um *userMiddleware) CreateUser(ctx *gin.Context) {
	logger.Info("[MIDDLEWARE (Create User)] Inicializando conexão com o banco de dados...")
	controller, database := um.handleConnection()
	defer mongodb.Disconnect(database)
	um.userController = controller
	um.userController.CreateUser(ctx)
}
