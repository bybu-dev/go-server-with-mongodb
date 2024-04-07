package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthUserResponse struct {
	ID string `json:"id"`
	AccountType string `json:"account_type"`
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
	AccountType string `json:"account_type"`

	Personal IPersonal `json:"personal"`
	Business []IBusiness `json:"business,omitempty"`
	Kyc IKyc `json:"kyc"`
	Setting ISetting `json:"setting"`
	Wallet IWallet `json:"wallet"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IGeneralUserResponse struct {
	ID primitive.ObjectID `json:"id"`
	AccountType string `json:"account_type"`

	Personal IPersonal `json:"personal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IMultipleSecureResponse struct {
	TotalUsers int64 `json:"total_users"`
	Users []ISecureUserResponse `json:"users"`
	HasNext bool `json:"has_next"`
}

type IMultipleGeneralResponse struct {
	TotalUsers int64 `json:"total_users"`
	Users []IGeneralUserResponse `json:"users"`
	HasNext bool `json:"has_next"`
}