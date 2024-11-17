package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // MongoDB auto-generated ID
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
