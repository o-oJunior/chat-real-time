package dto

import "time"

type UserDTO struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"createAt"`
	Password  string
}
