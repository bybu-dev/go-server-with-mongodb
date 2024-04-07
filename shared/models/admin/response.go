package admin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthUserResponse struct {
	ID string `json:"id"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type IPasswordAuthResponse struct {
	ID string `json:"id"`
	EmailAddress string `json:"account_type"`
	Message string `json:"message"`
}

type ISecureUserResponse struct {
	ID primitive.ObjectID `json:"id"`

	Personal IPersonal `json:"personal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IGeneralUserResponse struct {
	ID primitive.ObjectID `json:"id"`

	Personal IPersonal `json:"personal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}