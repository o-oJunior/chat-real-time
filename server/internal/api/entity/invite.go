package entity

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invite struct {
	ID                      string             `json:"id,omitempty" bson:"_id,omitempty"`
	Status                  string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAtAtMilliseconds int64              `json:"omitempty" bson:"createdAt,omitempty"`
	CreatedAt               string             `json:"createdAt,omitempty" bson:"omitempty"`
	UserIdInvited           primitive.ObjectID `json:"userIdInvited,omitempty" bson:"userIdInvited,omitempty"`
	UserIdInviter           primitive.ObjectID `json:"userIdInviter,omitempty" bson:"userIdInviter,omitempty"`
}

func (invite *Invite) ValidateRegisterInvite() error {
	if invite.UserIdInvited.Hex() == "" || invite.CreatedAt == "" || invite.Status == "" {
		return fmt.Errorf("o corpo da requisição está vazio ou mal formatado")
	}
	return nil
}
