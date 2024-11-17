package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Goal struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // MongoDB auto-generated ID
	UserID      primitive.ObjectID `bson:"user_id"`
	Target      float64            `bson:"target"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
}
