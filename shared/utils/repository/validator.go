package repo

import (
	"bybu/go-mongo-db/shared/models"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IValidatorRepository interface {
	Body(c *fiber.Ctx, request interface{}, duration time.Duration) models.IErrors
	Query(c *fiber.Ctx, request interface{}, duration time.Duration) models.IErrors
	Param(c *fiber.Ctx, request interface{}, duration time.Duration) models.IErrors
}

type ValidatorRepository struct {
	validate *validator.Validate
}

// Body implements IValidatorRepository.
func (repo *ValidatorRepository) Body(c *fiber.Ctx, request interface{}, duration time.Duration) models.IErrors {
	_, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	if err := c.BodyParser(request); err != nil {
		print(err.Error())
		return models.IErrors{models.IError{Message: "unable to decode this body"}}
	}

	if err := repo.validate.Struct(request); err != nil {
		var errors = models.IErrors{}
		for _, e := range err.(validator.ValidationErrors) {
			_error := models.IError{Field: e.Field(), Message: e.Tag()}
			errors = append(errors, _error)
		}
		return errors
	}

	return nil
}

// Param implements IValidatorRepository.
func (repo *ValidatorRepository) Param(c *fiber.Ctx, request interface{}, duration time.Duration) models.IErrors {
	_, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	if err := c.BodyParser(request); err != nil {
		print(err.Error())
		return models.IErrors{models.IError{Message: "unable to decode this body"}}
	}

	if err := repo.validate.Struct(request); err != nil {
		var errors = models.IErrors{}
		for _, e := range err.(validator.ValidationErrors) {
			_error := models.IError{Field: e.Field(), Message: e.Tag()}
			errors = append(errors, _error)
		}
		return errors
	}

	return nil
}

// Query implements IValidatorRepository.
func (repo *ValidatorRepository) Query(c *fiber.Ctx, request interface{}, duration time.Duration) models.IErrors {
	_, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	if err := c.BodyParser(request); err != nil {
		print(err.Error())
		return models.IErrors{models.IError{Message: "unable to decode this body"}}
	}

	if err := repo.validate.Struct(request); err != nil {
		var errors = models.IErrors{}
		for _, e := range err.(validator.ValidationErrors) {
			_error := models.IError{Field: e.Field(), Message: e.Tag()}
			errors = append(errors, _error)
		}
		return errors
	}

	return nil
}

func NewValidatorRepository() IValidatorRepository {
	return &ValidatorRepository{
		validate: validator.New(),
	}
}