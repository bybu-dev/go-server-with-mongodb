package authRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
	repo "bybu/go-mongo-db/shared/utils/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IAuthController struct {
	validator repo.IValidatorRepository
	service _IAuthService
}

func NewUserAuthController() *_IAuthController {
	return &_IAuthController{
		validator: repo.NewValidatorRepository(),
		service: *NewUserAuthService(),
	}
}

func (controller _IAuthController) register(c *fiber.Ctx) error {
	var request user.IRegisterRequest;

	errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.register(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller _IAuthController) login(c *fiber.Ctx) error {
	var request user.ILoginRequest

	errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.login(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller _IAuthController) accessToken(c *fiber.Ctx) error {
	request, _ := c.Locals("user").(user.ISecureUserResponse);
	
	errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.accessToken(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}