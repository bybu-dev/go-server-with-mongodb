package passwordRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"time"
)

type _IPasswordService struct {
	userRepo user.IUserRepository
}

func NewUserPasswordService() *_IPasswordService {
	return &_IPasswordService{ userRepo: user.NewUserRepository() }
}

func (service _IPasswordService) sendResetCode(request user.ISendResetCodeRequest) (user.IPasswordAuthResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: request.EmailAddress };
	updatedParam := user.IUserCodeInjection{ AuthenticationCode: "0000", Timeout: time.Now() };

	var userResponse, err = service.userRepo.Update( userParam, updatedParam);
	if (err != models.IError{}) {
		return user.IPasswordAuthResponse{}, models.IError{ 
			Field: "email_address",
			Message: "email address was not found",
		}
	}

	mailErr := utils.EmailSender.SendVerification("0000");
	if (mailErr != nil) {
		return user.IPasswordAuthResponse{}, models.IError{ 
			Field: "email_address",
			Message: "unable to send verification code",
		}
	}

	return user.IPasswordAuthResponse{
		ID: userResponse.ID.Hex(),
		EmailAddress: request.EmailAddress,
		Message: "code has been sent to this email",
	}, models.IError{}
}

func (service _IPasswordService) resetPassword(userRequest user.IVerifyResetCodeRequest) (user.IPasswordAuthResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }

	var userResponse user.IUser
	userResponse, responseErr := service.userRepo.Get(&userParam);
	if (responseErr != models.IError{}) {
		return user.IPasswordAuthResponse{}, models.IError{
			Field: "email_address",
			Message: "email is not registered with us",
		}
	}

	if (userResponse.Pheripheral.AuthenticationCode != userRequest.Code) {
		return user.IPasswordAuthResponse{}, models.IError{ Field: "email_address", Message: "invalid code"}
	}
	if (userResponse.Pheripheral.Timeout.Before(time.Now().Add(time.Hour*24*7))) {
		return user.IPasswordAuthResponse{}, models.IError{ Field: "email_address", Message: "code expired"}
	}

	return user.IPasswordAuthResponse{
		ID: userResponse.ID.String(),
		EmailAddress: userRequest.EmailAddress,
		Message: "code has been sent to this email",
	}, models.IError{}
}
