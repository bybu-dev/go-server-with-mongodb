package authRoutes

import (
	"bybu/go-mongo-db/shared/config"
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
	"bybu/go-postgres/shared/module/utils"
	"time"
)

type _IAuthService struct {
	userRepo user.IUserRepository
}

func NewUserAuthService() *_IAuthService{
	return &_IAuthService{
		userRepo: user.NewUserRepository(),
	}
}

func (service *_IAuthService) register(request user.IRegisterRequest) (user.IAuthUserResponse, models.IError) {
	hashPassword, err := utils.HashPassword(request.Password); if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "password is invalid" };
	}
	request.Password = hashPassword;
	newUser, userErr := service.userRepo.Create(request);
	if (userErr != models.IError{}) {
		return user.IAuthUserResponse{}, userErr
	}

	accessToken, tokenErr := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 48),
		Payload: newUser.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (tokenErr != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}
	
	refreshToken, tokenErr := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: newUser.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretRefreshKey(),
	});
	if (tokenErr != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create refresh token" }
	}

	return user.IAuthUserResponse{
		ID: newUser.ID.Hex(),
		AccountType: newUser.AccountType,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, models.IError{}
}

func (service *_IAuthService) login(userRequest user.ILoginRequest) (user.IAuthUserResponse, models.IError) {
	userParam := user.IPersonalRequest{ EmailAddress: userRequest.EmailAddress }
	userResponse, userErr := service.userRepo.Get(userParam);
	if (userErr != models.IError{}) {
		return user.IAuthUserResponse{}, models.IError{ Field: "password", Message: "invalid credencials" }
	}

	passwordErr := utils.CompareHashPassword(userRequest.Password, userResponse.Password);
	if (passwordErr != nil){
		return user.IAuthUserResponse{}, models.IError{ Field: "password", Message: "invalid credencials" }
	}

	accessToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 48),
		Payload: userResponse.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}
	refreshToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: userResponse.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretRefreshKey(),
	});
	if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}

	return user.IAuthUserResponse{
		ID: userResponse.ID.Hex(),
		AccountType: userResponse.AccountType,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, models.IError{}
}

func (service *_IAuthService) accessToken(userRequest user.ISecureUserResponse) (user.IAuthUserResponse, models.IError) {
	accessToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: userRequest.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (err != nil) {
		return user.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}

	return user.IAuthUserResponse{
		ID: userRequest.ID.Hex(),
		AccountType: userRequest.AccountType,
		AccessToken: accessToken,
	}, models.IError{}
}
