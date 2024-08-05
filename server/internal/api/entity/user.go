package entity

import (
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string    `json:"username" bson:"username"`
	FirstName    string    `json:"firstName" bson:"firstName"`
	LastName     string    `json:"lastName" bson:"lastName"`
	Email        string    `json:"email" bson:"email"`
	CreateAt     time.Time `json:"createAt" bson:"createAt"`
	Password     string    `json:"password,omitempty" bson:"password,omitempty"`
	HashPassword string    `json:"hashPassword,omitempty" bson:"hashPassword,omitempty"`
	Token        string    `json:"token,omitempty" bson:"token,omitempty"`
}

func errorParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) é obrigatório", name, typ)
}

func (user *User) ValidateCreateUser() error {
	if user.Username == "" && user.FirstName == "" &&
		user.LastName == "" && user.Email == "" && user.Password == "" {
		return fmt.Errorf("o corpo da requisição está vazio ou mal formatado")
	}

	if user.Username != "" {
		validNameRegex := "^[a-zA-Z0-9]+$"
		matched, _ := regexp.MatchString(validNameRegex, user.Username)
		if !matched {
			return fmt.Errorf("username não pode conter caracteres especiais")
		}
	}

	if user.Username == "" {
		return errorParamIsRequired("username", "string")
	}

	if user.FirstName == "" {
		return errorParamIsRequired("firstName", "string")
	}

	if user.LastName == "" {
		return errorParamIsRequired("lastName", "string")
	}

	if user.Email == "" {
		return errorParamIsRequired("email", "string")
	}

	if user.Password == "" {
		return errorParamIsRequired("password", "string")
	}

	return nil
}

func (user *User) EncodePassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	user.HashPassword = string(hashPassword)
	user.Password = ""
	return nil
}

func (user *User) ComparePassword(hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return fmt.Errorf("senha inválida")
	}
	return nil
}

func (user *User) GetID() string {
	return user.ID
}
