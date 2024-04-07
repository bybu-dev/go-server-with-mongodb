package walletRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
	repo "bybu/go-mongo-db/shared/utils/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IProfileController struct {
	validator repo.IValidatorRepository
	service _IProfileService
}

func NewUserWalletController() *_IProfileController {
	return &_IProfileController{
		validator: repo.NewValidatorRepository(),
		service: *NewUserWalletService(),
	}
}

func (controller _IProfileController) fund(c *fiber.Ctx) error {
	var request user.IPersonal;
	err := controller.validator.Body(c, request, time.Duration(2*time.Second)); 
	if (err != nil) {
		c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(err))
	}

	response, errResponse := controller.service.fund(c.Locals("user").(user.IUser), request);
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}

func (controller _IProfileController) withdraw(c *fiber.Ctx) error {
	var request user.IPersonal;
	err := controller.validator.Body(c, request, time.Duration(2*time.Second)); 
	if (err != nil) {
		c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(err));
	}

	response, errResponse := controller.service.withdraw(c.Locals("user").(user.IUser), request);
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}