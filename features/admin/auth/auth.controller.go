package adminAuthRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/admin"
	repo "bybu/go-mongo-db/shared/utils/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IAuthController struct {
	service _IAuthService
	validator repo.IValidatorRepository
}

func NewAdminAuthController() *_IAuthController {
	return &_IAuthController{
		service: *NewAdminAuthService(),
		validator: repo.NewValidatorRepository(),
	}
}

func (ac _IAuthController) create(c *fiber.Ctx) error {
	var request admin.ICreateRequest;

	errors := ac.validator.Body(c, &request, time.Duration(2*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	} 

	response, err := ac.service.create(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err }));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (ac _IAuthController) login(c *fiber.Ctx) error {
	var request admin.ILoginRequest

	errors := ac.validator.Body(c, &request, time.Duration(2*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := ac.service.login(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err },));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}