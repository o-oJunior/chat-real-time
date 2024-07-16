package middleware

import (
	"server/src/config"
	"server/src/config/mongodb"
	"server/src/controller"
	"server/src/model/repository"
	"server/src/model/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var log *config.Logger = config.NewLogger("middleware")

type UserMiddleware interface {
	CreateUser(*gin.Context)
}

type userMiddleware struct {
	userController controller.UserController
}

func NewUserMiddleware() UserMiddleware {
	return &userMiddleware{}
}

func (userMiddleware *userMiddleware) handleConnection() (controller.UserController, *mongo.Database) {
	database := mongodb.Connect()
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController, database
}

func (middleware *userMiddleware) CreateUser(ctx *gin.Context) {
	log.Info("(Create User) Inicializando conexão com o banco de dados...")
	controller, database := middleware.handleConnection()
	defer mongodb.Disconnect(database)
	middleware.userController = controller
	middleware.userController.CreateUser(ctx)
}
