package user

import (
	"time"
)

type IRegisterRequest struct {
	Personal IPersonal `json:"personal"`
	Password string `json:"password"`
}

type ILoginRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ICreateTokenRequest struct {
	UserID string `json:"u" bson:"_id" validate:"required"`
}

type ISendResetCodeRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type IVerifyResetCodeRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Code string `json:"auth_code" validate:"required"`
	Password string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type IPersonalRequest struct {
	EmailAddress string `json:"email_address" bson:"personal.email_address"`
}

type IUpdatePersonal struct {
	FirstName string `json:"first_name," bson:"first_name" validate:"required"`
	SecondName string `json:"second_name" bson:"second_name" validate:"required"`
	EmailAddress string `json:"email_address" bson:"email_address,unique" validate:"required,email"`
	ProfileImage string `json:"profile_image" bson:"profile_image"`
	Username string `json:"username" bson:"username"`
}

type IUserCodeInjection struct {
	AuthenticationCode string `json:"authentication_code" bson:"pheripheral.authentication_code"`
	Timeout time.Time `json:"timeout" bson:"pheripheral.timeout"`
	IsBanned bool `json:"is_banned" bson:"is_banned,false"`
}