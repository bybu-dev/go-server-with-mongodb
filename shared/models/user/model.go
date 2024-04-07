package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IWallet struct {
	Balance int `json:"balance" bson:"balance"`
	TotalBalance int `json:"total_balance" bson:"total_balance"`
}

type IKyc struct {
	VerificationStatus string `json:"verification_status" bson:"verification_status"`
	VerificationDocument []string `json:"verification_document" bson:"verification_document"`
}

type IPersonal struct {
	FirstName string `json:"first_name" bson:"first_name" validate:"required"`
	SecondName string `json:"second_name" bson:"second_name" validate:"required"`
	EmailAddress string `json:"email_address" bson:"email_address,unique" validate:"required,email"`
	ProfileImage string `json:"profile_image" bson:"profile_image"`
	Username string `json:"username" bson:"username"`
}

type IPheripheral struct {
	AuthenticationCode string `json:"authentication_code" bson:"authentication_code"`
	Timeout time.Time `json:"timeout" bson:"timeout"`
	IsBanned bool `json:"is_banned" bson:"is_banned,false"`
}

type ISetting struct {
	EmailReminder bool `json:"email_reminder" bson:"email_reminder"`
	PopupNotification bool `json:"pop_up_notification" bson:"pop_up_notification"`
	IsInformationEditable bool `json:"is_information_editable" bson:"is_information_editable"`
}

type IBusiness struct {
	BusinessName string `json:"business_name" bson:"business_name"`
	ContactPhoneNumber string `json:"contact_phone_number" bson:"contact_phone_number"`
	BusinessAddress string `json:"business_address" bson:"business_address"`
	BusinessType string `json:"business_type" bson:"business_type"`
}

type IUser struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AccountType string `json:"account_type" bson:"account_type,omitempty"`

	Personal IPersonal `json:"personal" bson:"personal,omitempty"`
	Kyc IKyc `json:"kyc" bson:"kyc"`
	Setting ISetting `json:"setting" bson:"setting,omitempty"`
	Wallet IWallet `json:"wallet" bson:"wallet,omitempty"`
	Pheripheral IPheripheral `json:"pheripheral" bson:"pheripheral,omitempty"`
	Business []IBusiness `json:"business_information,omitempty" bson:"business_information,omitempty"`

	Password string `json:"password" bson:"password,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

type IMultipleUser struct {
	TotalUsers int64 `json:"total_users"`
	Users []IUser `json:"users"`
	HasNext bool `json:"has_next"`
}


type MultipleUser interface {
	int64
	[]IUser | []IGeneralUserResponse | []ISecureUserResponse
	bool
}


