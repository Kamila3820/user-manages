package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Credential struct {
		Id           primitive.ObjectID `bson:"_id,omitempty"`
		UserId       string             `bson:"user_id"`
		AccessToken  string             `bson:"access_token"`
		RefreshToken string             `bson:"refresh_token"`
		CreatedAt    time.Time          `bson:"created_at"`
		UpdatedAt    time.Time          `bson:"updated_at"`
	}

	Role struct {
		Id    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title string             `json:"title" bson:"title"`
		Code  int                `json:"code" bson:"code"`
	}

	UpdateRefreshTokenReq struct {
		UserId       string    `bson:"user_id"`
		AccessToken  string    `bson:"access_token"`
		RefreshToken string    `bson:"refresh_token"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
)
