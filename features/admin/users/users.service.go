package adminUsersRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
)

type _IUsersService struct {
	userRepo user.IUserRepository
}

func newUserService() *_IUsersService {
	return &_IUsersService{
		userRepo: user.NewUserRepository(),
	}
}

func (service *_IUsersService) getAllUsers(request models.IOptions) (user.IMultipleGeneralResponse, models.IError) {
	multipleUsers, userErr := service.userRepo.GetMultiple(user.IUser{}, request);
	if (userErr != models.IError{}) {
		return user.IMultipleGeneralResponse{}, userErr
	}

	return multipleUsers.ToGeneralResponse(), models.IError{}
}

func (service *_IUsersService) banUser(userRequest user.IPersonalRequest) (user.IGeneralUserResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }

	userResponse, userErr := service.userRepo.Update(userParam, user.IUserCodeInjection{IsBanned: true});
	if (userErr != models.IError{}) {
		return user.IGeneralUserResponse{}, models.IError{ Message: "unable to ban this user" }
	}

	return userResponse.ToGeneralResponse(), models.IError{}
}

func (service *_IUsersService) unbanUser(userRequest user.IPersonalRequest) (user.IGeneralUserResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }

	userResponse, userErr := service.userRepo.Update(userParam, user.IUserCodeInjection{IsBanned: false});
	if (userErr != models.IError{}) {
		return user.IGeneralUserResponse{}, models.IError{ Message: "unable to ban this user" }
	}

	return userResponse.ToGeneralResponse(), models.IError{}
}