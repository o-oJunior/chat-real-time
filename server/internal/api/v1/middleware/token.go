package middleware

import (
	"fmt"
	"net/http"
	"os"
	"server/internal/api/entity"
	"server/internal/api/v1/response"
	"server/internal/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

type Token interface {
	Generate(user *entity.User) (string, error)
	ValidateCookie(ctx *gin.Context)
}

type token struct {
	PRIVATE_KEY []byte
}

type customClaims struct {
	ID          string `json:"id,omitempty"`
	Username    string `json:"username"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreateAt    string `json:"createAt"`
	jwt.StandardClaims
}

func NewMiddlewareToken() Token {
	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")
	return &token{[]byte(PRIVATE_KEY)}
}

var logger *config.Logger = config.NewLogger("middleware token")

func (tkn token) Generate(user *entity.User) (string, error) {
	claims := customClaims{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Status:      user.Status,
		Description: user.Description,
		CreateAt:    user.CreateAt,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tkn.PRIVATE_KEY)
	if err != nil {
		return "", fmt.Errorf("erro a gerar o token")
	}
	return tokenString, nil
}

func (tkn token) ValidateCookie(ctx *gin.Context) {
	logger.Info("Conferindo se token é válido...")
	cookie, err := ctx.Cookie("token")
	if err != nil {
		logger.Error("Token não veio nos cookies")
		response.SendError(ctx, http.StatusUnauthorized, "Token não veio nos cookies")
		ctx.Abort()
		return
	}
	token, err := jwt.Parse(cookie, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", jwtToken.Header["alg"])
		}
		return tkn.PRIVATE_KEY, nil
	})

	if err != nil {
		logger.Error("Token expirado")
		response.SendError(ctx, http.StatusUnauthorized, "Token expirado")
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Error("Token inválido")
		response.SendError(ctx, http.StatusUnauthorized, "Token inválido")
		ctx.Abort()
		return
	}
	logger.Info("Token válido, será retornado os dados do usuário")
	data := bson.M{
		"id":          claims["id"],
		"username":    claims["username"],
		"firstName":   claims["firstName"],
		"lastName":    claims["lastName"],
		"email":       claims["email"],
		"status":      claims["status"],
		"description": claims["description"],
		"createAt":    claims["createAt"],
		"expiresAt":   claims["exp"],
	}
	response.SendSuccess(ctx, http.StatusOK, "", data)
}
