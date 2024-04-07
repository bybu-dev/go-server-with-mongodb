package user

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (user IUser) ToGeneralResponse() IGeneralUserResponse {
	generalUser := IGeneralUserResponse{
		ID: user.ID,
		AccountType: user.AccountType,
	
		Personal: user.Personal,
	
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return generalUser
}
	
func(user IUser) ToSecureResponse() ISecureUserResponse {
	secureUser := ISecureUserResponse{
		ID: user.ID,
		AccountType: user.AccountType,
	
		Personal: user.Personal,
		Business: user.Business,
		Kyc: user.Kyc,
		Setting: user.Setting,
		Wallet: user.Wallet,
	
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return secureUser
}

func (request IRegisterRequest) CreateUser() IUser {
	return IUser{
		ID: primitive.NewObjectID(),
		AccountType: "Personal",
		Personal: IPersonal{
			FirstName: request.Personal.FirstName,
			SecondName: request.Personal.SecondName,
			EmailAddress: strings.ToLower(request.Personal.EmailAddress),
		},
		Kyc: IKyc{
			VerificationStatus: "unverified",
		},
		Setting: ISetting{
			EmailReminder: false,
			PopupNotification: false,
			IsInformationEditable: true,
		},
		Password: request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (user IMultipleUser) ToGeneralResponse() IMultipleGeneralResponse {
	var users []IGeneralUserResponse;
	for i := 0; i < len(user.Users); i++ {
		users = append(users, user.Users[i].ToGeneralResponse());
	}
	generalUser := IMultipleGeneralResponse{
		TotalUsers: user.TotalUsers,
		Users: users,
		HasNext: user.HasNext,
	};
	return generalUser
}
	
func(user IMultipleUser) ToSecureResponse() IMultipleSecureResponse {
	var users []ISecureUserResponse;
	for i := 0; i < len(user.Users); i++ {
		users = append(users, user.Users[i].ToSecureResponse());
	}
	secureUser := IMultipleSecureResponse{
		TotalUsers: user.TotalUsers,
		Users: users,
		HasNext: user.HasNext,
	};
	return secureUser
}