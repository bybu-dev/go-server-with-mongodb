package profileRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
)

type _IProfileService struct {
	userRepo user.IUserRepository
}

func NewUserProfileService() *_IProfileService {
	return &_IProfileService{ userRepo: user.NewUserRepository() }
}

func (service *_IProfileService) getProfile(me user.IUser) (user.ISecureUserResponse, models.IError) {

	return me.ToSecureResponse(), models.IError{};
}

func (service *_IProfileService) updateProfile(me user.IUser, update user.IPersonal) (user.ISecureUserResponse, models.IError) {
	request := user.IUser{ ID: me.ID };
	updatedUser, err := service.userRepo.Update(request, update);
	if (err != models.IError{}) {
		return user.ISecureUserResponse{}, err;
	}

	return updatedUser.ToSecureResponse(), models.IError{};
}