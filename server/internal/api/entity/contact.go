package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	ID                    string             `json:"id,omitempty" bson:"_id,omitempty"`
	InviteStatus          string             `json:"inviteStatus,omitempty" bson:"inviteStatus,omitempty"`
	InvitedAtMilliseconds int64              `json:"omitempty" bson:"invitedAt,omitempty"`
	UserIdInvited         primitive.ObjectID `json:"userIdInvited,omitempty" bson:"userIdInvited,omitempty"`
	UserIdInviter         primitive.ObjectID `json:"userIdInviter,omitempty" bson:"userIdInviter,omitempty"`
}
