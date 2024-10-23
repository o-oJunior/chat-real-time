package middleware

import (
	"fmt"
	"net/http"
	"server/internal/api/entity"
	"server/internal/api/v1/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token interface {
	DecodeToken(string) (primitive.M, error)
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
	CreatedAt   string `json:"createdAt"`
	jwt.StandardClaims
}

func (tkn token) Generate(user *entity.User) (string, error) {
	claims := customClaims{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Status:      user.Status,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
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
	logger.Info("(Token) Conferindo se token é válido...")
	cookie, err := ctx.Cookie("token")
	if err != nil {
		logger.Error("(Token) Token não veio nos cookies")
		response.SendError(ctx, http.StatusUnauthorized, "Token não veio nos cookies")
		ctx.Abort()
		return
	}
	data, err := tkn.DecodeToken(cookie)
	if err != nil {
		logger.Error("Erro ao decifrar o token: %v", err)
		ctx.Abort()
		return
	}
	response.SendSuccess(ctx, http.StatusOK, "", data)
}

func (tkn token) DecodeToken(cookie string) (primitive.M, error) {
	token, err := jwt.Parse(cookie, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("(Token) método de assinatura inesperado: %v", jwtToken.Header["alg"])
		}
		return tkn.PRIVATE_KEY, nil
	})

	if err != nil {
		logger.Error("(Token) Token expirado")
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Error("(Token) Token inválido")
		return nil, err
	}
	logger.Info("(Token) Token válido, será retornado os dados do usuário")
	data := bson.M{
		"id":          claims["id"],
		"username":    claims["username"],
		"firstName":   claims["firstName"],
		"lastName":    claims["lastName"],
		"email":       claims["email"],
		"status":      claims["status"],
		"description": claims["description"],
		"createdAt":   claims["createdAt"],
		"expiresAt":   claims["exp"],
	}
	return data, nil
}
