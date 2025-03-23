package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Message      string             `json:"message,omitempty" bson:"message,omityempty"`
	UserIdTarget primitive.ObjectID `json:"userIdTarget,omitempty" bson:"userIdTarget,omitempty"`
	UserIdActor  primitive.ObjectID `json:"userIdActor,omitempty" bson:"userIdActor,omitempty"`
	Type         string             `json:"type,omitempty" bson:"type,omityempty"`
	CreatedAt    int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
