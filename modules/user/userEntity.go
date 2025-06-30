package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email     string             `json:"email" bson:"email"`
		Password  string             `json:"password" bson:"password"`
		Name      string             `json:"name" bson:"name"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	}

	UserProfileBson struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email     string             `json:"email" bson:"email"`
		Name      string             `json:"name" bson:"name"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
