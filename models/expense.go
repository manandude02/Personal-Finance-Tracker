package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Expense struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // MongoDB auto-generated ID
	UserID      primitive.ObjectID `bson:"user_id"`
	Category    string             `bson:"category"`
	Amount      float64            `bson:"amount"`
	Date        time.Time          `bson:"date"`
	Description string             `bson:"description"`
}
