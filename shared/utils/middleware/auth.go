package middleware

import (
	"bybu/go-mongo-db/shared/config"
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/models/admin"
	"bybu/go-mongo-db/shared/models/user"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthenticate struct {
	userRepo user.IUserRepository
	adminRepo admin.IAdminRepository
}


var Validate = IAuthenticate{
	adminRepo: admin.NewAdminRepository(),
	userRepo: user.NewUserRepository(),
}

func (ur *IAuthenticate) UserRole(c *fiber.Ctx) error {
	_, ctxFunc := context.WithTimeout(context.Background(), time.Second * 2)
	defer ctxFunc()
	var tokenString string = "";
	authorization := c.Get("Authorization");

	if (strings.HasPrefix(authorization, "Bearer ")) {
		tokenString = strings.TrimPrefix(authorization, "Bearer ");
	} else {
		tokenString = c.Cookies("token");
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "You are not logged in" },
		}))
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.Env.GetUserSecretKey()), nil
	})
	if (err != nil) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "invalidate token: " + err.Error(),
		}}));
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "invalid token claim" },
		}));

	}

	objectID, objectErr := primitive.ObjectIDFromHex((claims["sub"]).(string))
	if (objectErr != nil) {
		c.Status(401).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "Unable to decode user ID"},
		}))
	}

	var userParam = user.IUser{ ID: objectID }
	user, userErr := ur.userRepo.Get(userParam);
	if (userErr != models.IError{}) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "unable to get this user" },
		}));
	}

	c.Locals("user", user);

	return c.Next()
}

func (ur *IAuthenticate) RefreshRole(c *fiber.Ctx) error {
	_, ctxFunc := context.WithTimeout(context.Background(), time.Second * 2)
	ctxFunc()
	var tokenString string = "";
	authorization := c.Get("Authorization");

	if (strings.HasPrefix(authorization, "Bearer ")) {
		tokenString = strings.TrimPrefix(authorization, "Bearer ");
	} else {
		tokenString = c.Cookies("token");
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "You are not logged in" },
		}))
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.Env.GetUserSecretRefreshKey()), nil
	})
	if (err != nil) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "invalidate token: " + err.Error(),
		}}));
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "invalid token claim" },
		}));

	}

	objectID, objectErr := primitive.ObjectIDFromHex((claims["sub"]).(string))
	if (objectErr != nil) {
		c.Status(401).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "Unable to decode user ID"},
		}))
	}

	var userParam = user.IUser{ ID: objectID }
	user, userErr := ur.userRepo.Get(userParam);
	if (userErr != models.IError{}) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "unable to get this user" },
		}));
	}

	c.Locals("user", user);

	return c.Next()
}

func (ar *IAuthenticate) AdminRole(c *fiber.Ctx) error {
	_, ctxFunc := context.WithTimeout(context.Background(), time.Second * 2)
	defer ctxFunc()
	var tokenString string = "";
	authorization := c.Get("Authorization");

	if (strings.HasPrefix(authorization, "Bearer ")) {
		tokenString = strings.TrimPrefix(authorization, "Bearer ");
	} else {
		tokenString = c.Cookies("token");
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "You are not logged in" },
		}))
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(config.Env.GetUserSecretKey()), nil
	})
	if (err != nil) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "invalidate token: " + err.Error(),
		}}));
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "invalid token claim" },
		}));

	}

	objectID, objectErr := primitive.ObjectIDFromHex((claims["sub"]).(string))
	if (objectErr != nil) {
		c.Status(401).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "Unable to decode user ID"},
		}))
	}

	var userParam = user.IUser{ ID: objectID }
	user, userErr := ar.adminRepo.Get(userParam);
	if (userErr != models.IError{}) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ToErrorResponse(models.IErrors{
			models.IError{ Message: "unable to get this user" },
		}));
	}

	c.Locals("user", user);

	return c.Next()
}