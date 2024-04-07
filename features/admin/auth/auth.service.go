package adminAuthRoutes

import (
	"bybu/go-mongo-db/shared/config"
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/admin"
	"bybu/go-postgres/shared/module/utils"
	"time"
)

type _IAuthService struct {
	admin admin.IAdminRepository
}

func NewAdminAuthService() *_IAuthService {
	return &_IAuthService{
		admin: admin.NewAdminRepository(),
	}
};

func (service _IAuthService) create(request admin.ICreateRequest) (admin.IAuthUserResponse, models.IError) {
	hashPassword, err := utils.HashPassword(request.Password); if (err != nil) {
		return admin.IAuthUserResponse{}, models.IError{ Message: "password is invalid" };
	}
	request.Password = hashPassword;
	newUser, adminErr := service.admin.Create(request);
	if (adminErr != models.IError{}) {
		return admin.IAuthUserResponse{}, adminErr
	}

	accessToken, tokenErr := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 48),
		Payload: newUser.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (tokenErr != nil) {
		return admin.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}
	
	refreshToken, tokenErr := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: newUser.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretRefreshKey(),
	});
	if (tokenErr != nil) {
		return admin.IAuthUserResponse{}, models.IError{ Message: "unable to create refresh token" }
	}

	return admin.IAuthUserResponse{
		ID: newUser.ID.Hex(),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, models.IError{}
}

func (service _IAuthService) login(adminRequest admin.ILoginRequest) (admin.IAuthUserResponse, models.IError) {
	adminParam := admin.IPersonalRequest{ EmailAddress: adminRequest.EmailAddress }
	adminResponse, adminErr := service.admin.Get(adminParam);
	if (adminErr != models.IError{}) {
		return admin.IAuthUserResponse{}, models.IError{ Field: "password", Message: "invalid credencials" }
	}

	passwordErr := utils.CompareHashPassword(adminRequest.Password, adminResponse.Password);
	if (passwordErr != nil){
		return admin.IAuthUserResponse{}, models.IError{ Field: "password", Message: "invalid credencials" }
	}

	accessToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 48),
		Payload: adminResponse.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretKey(),
	});
	if (err != nil) {
		return admin.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}
	refreshToken, err := utils.CreateToken(utils.TokenParams{
		Ttl: time.Now().Add(time.Hour * 24 * 180),
		Payload: adminResponse.ID.Hex(),
		PrivateKey: config.Env.GetUserSecretRefreshKey(),
	});
	if (err != nil) {
		return admin.IAuthUserResponse{}, models.IError{ Message: "unable to create access token" }
	}

	return admin.IAuthUserResponse{
		ID: adminResponse.ID.Hex(),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, models.IError{}
}
