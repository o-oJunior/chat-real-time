package entity

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   string `json:"id,omitempty" bson:"_id,omitempty"`
	Username             string `json:"username" bson:"username"`
	FirstName            string `json:"firstName" bson:"firstName"`
	LastName             string `json:"lastName" bson:"lastName"`
	Email                string `json:"email" bson:"email"`
	CreateAtMilliseconds int64  `bson:"createAt"`
	CreateAt             string `json:"createAt"`
	Password             string `json:"password,omitempty" bson:"password,omitempty"`
	HashPassword         string `json:"hashPassword,omitempty" bson:"hashPassword,omitempty"`
	Token                string `json:"token,omitempty" bson:"token,omitempty"`
}

func errorParamIsRequired(name string) error {
	return fmt.Errorf("%s é obrigatório", name)
}

func errorParamMinimunValue(name string, length int8) error {
	return fmt.Errorf("%s deve conter no minimo %d caracteres", name, length)
}

func errorParamMaximunValue(name string, length int8) error {
	return fmt.Errorf("%s deve conter no máximo %d caracteres", name, length)
}

func validateEmail(email string) bool {
	var regexEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return regexEmail.MatchString(email)
}

func (user *User) ValidateCreateUser() error {
	if user.Username == "" && user.FirstName == "" &&
		user.LastName == "" && user.Email == "" && user.Password == "" {
		return fmt.Errorf("o corpo da requisição está vazio ou mal formatado")
	}

	fields := map[string]struct {
		value     string
		minLength int8
		maxLength int8
	}{
		"nome de usuário": {user.Username, 3, 20},
		"nome":            {user.FirstName, 3, 20},
		"sobrenome":       {user.LastName, 3, 50},
		"email":           {user.Email, 5, 40},
		"senha":           {user.Password, 5, 20},
	}

	for fieldName, field := range fields {
		if field.value == "" {
			return errorParamIsRequired(fieldName)
		} else if len(field.value) < int(field.minLength) {
			return errorParamMinimunValue(fieldName, field.minLength)
		} else if len(field.value) > int(field.maxLength) {
			return errorParamMaximunValue(fieldName, field.maxLength)
		}
		if fieldName == "nome de usuário" && field.value != "" {
			validRegex := "^[a-zA-Z0-9_^~`´@]+$"
			regex, _ := regexp.Compile(validRegex)
			matched := regex.MatchString(field.value)
			if !matched {
				return fmt.Errorf("%s contém caracteres especiais ou espaços em branco", fieldName)
			}
		} else if fieldName == "email" {
			isEmail := validateEmail(user.Email)
			if !isEmail {
				return fmt.Errorf("%s inválido", fieldName)
			}
		}
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
