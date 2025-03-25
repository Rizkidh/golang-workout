package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Name      string             `bson:"name" json:"name"`
	URL       string             `bson:"url" json:"url"`
	AIEnabled bool               `bson:"ai_enabled" json:"ai_enabled"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}
