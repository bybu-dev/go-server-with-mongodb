package admin

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (user IAdmin) ToUserResponse() ISecureUserResponse {
		secureUser := ISecureUserResponse{
			ID: user.ID,
			Personal: user.Personal,

			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		return secureUser
}

func (request ICreateRequest) CreateUser() IAdmin {
	return IAdmin{
		ID: primitive.NewObjectID(),
		Personal: IPersonal{
			FirstName: request.Personal.FirstName,
			SecondName: request.Personal.SecondName,
			EmailAddress: strings.ToLower(request.Personal.EmailAddress),
		},
		Password: request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}