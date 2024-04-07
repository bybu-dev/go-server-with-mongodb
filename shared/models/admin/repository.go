package admin

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/utils"
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAdminRepository interface {
	Create(request ICreateRequest) (IAdmin, models.IError)
	Get(request interface{}) (IAdmin, models.IError)
	Update(user interface{}, update interface{}) (IAdmin, models.IError)
}

type AdminMongoRepository struct {
	repo *mongo.Collection
}

func NewAdminRepository() IAdminRepository {
	return &AdminMongoRepository{
		repo: utils.GetCollection(utils.DB, "admins"),
	}
}

func (rp *AdminMongoRepository) Create(request ICreateRequest) (IAdmin, models.IError) {
	ctx := context.TODO()
	newUser := request.CreateUser();

	response, insertErr := rp.repo.InsertOne(ctx, &newUser);
	if (insertErr != nil) {
		if (strings.Contains(insertErr.Error(), "E11000")) {
			return IAdmin{}, models.IError{ Field: "email_address", Message: "user with this email address already exist" }
		}
		print(insertErr.Error());
		return IAdmin{}, models.IError{ Message: insertErr.Error() }
	}
	return IAdmin{ ID: response.InsertedID.(primitive.ObjectID) }, models.IError{}
}

func (rp *AdminMongoRepository) Get(request interface{}) (IAdmin, models.IError) {
	ctx := context.TODO()
	var userResponse IAdmin

	err := rp.repo.FindOne(ctx, &request).Decode(&userResponse);
	if (err != nil) {
		return IAdmin{}, models.IError{ Message: "unable to get this user" }
	}

	return userResponse, models.IError{};
}

func (rp *AdminMongoRepository) Update(user interface{}, update interface{}) (IAdmin, models.IError) {
	ctx := context.TODO()

	var userResponse, err = rp.repo.UpdateOne(ctx, user, update);
	if (err != nil) {
		return IAdmin{}, models.IError{ Message: "unable to get this user" }
	}

	return IAdmin{
		ID: userResponse.UpsertedID.(primitive.ObjectID),
	}, models.IError{};
}
