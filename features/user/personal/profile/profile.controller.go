package profileRoutes

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/user"
	repo "bybu/go-mongo-db/shared/utils/repository"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type _IProfileController struct {
	validator repo.IValidatorRepository
	service _IProfileService
}

func NewProfileController() *_IProfileController {
	return &_IProfileController{
		validator: repo.NewValidatorRepository(),
		service: *NewUserProfileService(),
	}
}

func (controller _IProfileController) getUser(c *fiber.Ctx) error {
	_, ctxErr := context.WithTimeout(context.Background(), time.Second);
	defer ctxErr();

	response, errResponse := controller.service.getProfile(c.Locals("user").(user.IUser));
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}

func (controller _IProfileController) updateUser(c *fiber.Ctx) error {
	var request user.IPersonal;
	err := controller.validator.Body(c, request, time.Duration(10*time.Second)); 
	if (err != nil) {
		c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(err))
	}

	response, errResponse := controller.service.updateProfile(c.Locals("user").(user.IUser), request);
	if (errResponse != models.IError{}) {
		return c.Status(http.StatusBadRequest).JSON(models.ToErrorResponse(models.IErrors{errResponse}))
	}

	return c.Status(http.StatusOK).JSON(models.ToSuccessResponse(response))
}