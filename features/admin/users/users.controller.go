package adminUsersRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
	repo "bybu/go-mongo-db/shared/utils/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IUserManagementController struct {
	service _IUsersService
	validator repo.IValidatorRepository
}

func newUserController() *_IUserManagementController {
	return &_IUserManagementController{
		service: *newUserService(),
		validator: repo.NewValidatorRepository(),
	}
}

func (controller *_IUserManagementController) getAllUsers(c *fiber.Ctx) error {
		var request models.IOptions;

		errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
		if (errors != nil) {
			return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
		}

		response, err := controller.service.getAllUsers(request);
		if (err != models.IError{}) {
			return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err }));
		}

		return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller *_IUserManagementController) banUser(c *fiber.Ctx) error {
	var request user.IPersonalRequest

	errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.banUser(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err },));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}

func (controller *_IUserManagementController) unbanUser(c *fiber.Ctx) error {
	var request user.IPersonalRequest

	errors := controller.validator.Body(c, &request, time.Duration(10*time.Second));
	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(errors));
	}

	response, err := controller.service.unbanUser(request);
	if (err != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{ err },));
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response));
}
