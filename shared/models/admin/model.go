package admin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
}

type IAdmin struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Personal IPersonal `json:"personal" bson:"personal,omitempty"`
	Pheripheral IPheripheral `json:"pheripheral" bson:"pheripheral,omitempty"`

	Password string `json:"password" bson:"password,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}