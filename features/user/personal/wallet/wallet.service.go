package walletRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
)

type _IProfileService struct {
	userRepo user.IUserRepository
}

func NewUserWalletService() *_IProfileService {
	return &_IProfileService{ userRepo: user.NewUserRepository() }
}

func (service *_IProfileService) fund(me user.IUser, update user.IPersonal) (user.IUser, models.IError) {
	request := user.IUser{ ID: me.ID };
	updatedUser, err := service.userRepo.Update(request, update);
	if (err != models.IError{}) {
		return user.IUser{}, err;
	}

	return updatedUser, models.IError{};
}

func (service *_IProfileService) withdraw(me user.IUser, update user.IPersonal) (user.IUser, models.IError) {
	request := user.IUser{ ID: me.ID };
	updatedUser, err := service.userRepo.Update(request, update);
	if (err != models.IError{}) {
		return user.IUser{}, err;
	}

	return updatedUser, models.IError{};
}