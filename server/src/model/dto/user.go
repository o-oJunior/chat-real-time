package dto

import "time"

type User struct {
	Username     string    `json:"username" bson:"username"`
	FirstName    string    `json:"firstName" bson:"firstName"`
	LastName     string    `json:"lastName" bson:"lastName"`
	Email        string    `json:"email" bson:"email"`
	CreateAt     time.Time `json:"createAt" bson:"createAt"`
	UpdateAt     time.Time `json:"updateAt" bson:"updateAt"`
	Password     string    `json:"password,omitempty" bson:"password,omitempty"`
	HashPassword string    `json:"hashPassword,omitempty" bson:"hashPassword,omitempty"`
}
