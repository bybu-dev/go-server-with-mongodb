package passwordRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
	repo "bybu/go-mongo-db/shared/utils/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IPasswordController struct {
	validator repo.IValidatorRepository
	service _IPasswordService
}

func NewUserPasswordController() *_IPasswordController {
	return &_IPasswordController{
		validator: repo.NewValidatorRepository(),
		service: *NewUserPasswordService(),
	}
}

func (controller _IPasswordController) sendResetCode(c *fiber.Ctx) error {
	var request user.ISendResetCodeRequest
	
	errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.sendResetCode(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller _IPasswordController) resetPassword(c *fiber.Ctx) error {
	var request user.IVerifyResetCodeRequest

	errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.resetPassword(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(
			models.IErrors{ err },
		));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}