package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
	UserIdTarget primitive.ObjectID `json:"userIdTarget,omitempty" bson:"userIdTarget,omitempty"`
	UserIdActor  primitive.ObjectID `json:"userIdActor,omitempty" bson:"userIdActor,omitempty"`
	CreatedAt    int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    int64              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
