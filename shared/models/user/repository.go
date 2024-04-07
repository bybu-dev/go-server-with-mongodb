package user

import (
	"bybu/go-mongo-db/shared/models"
	"bybu/go-mongo-db/shared/utils"
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var UserCollection *mongo.Collection = module.GetCollection(module.DB, "users");

type IUserRepository interface {
	Create(request IRegisterRequest) (IUser, models.IError)
	Get(request interface{}) (IUser, models.IError)
	GetMultiple(search IUser, option models.IOptions)(IMultipleUser, models.IError)
	Update(user interface{}, update interface{}) (IUser, models.IError)
}

type MongoUserRepository struct {
	repo *mongo.Collection
}

func NewUserRepository() IUserRepository {
	return &MongoUserRepository{ 
		repo: utils.GetCollection(utils.DB, "users"),
	};
}

func (ur *MongoUserRepository) Create(request IRegisterRequest) (IUser, models.IError) {
	ctx := context.TODO();
	newUser := request.CreateUser();

	response, insertErr := ur.repo.InsertOne(ctx, &newUser);
	if (insertErr != nil) {
		if (strings.Contains(insertErr.Error(), "E11000")) {
			return IUser{}, models.IError{ Field: "email_address", Message: "user with this email address already exist" }
		}
		print("Freedom2: ");
		print(insertErr.Error());
		return IUser{}, models.IError{ Message: insertErr.Error() }
	}
	return IUser{ ID: response.InsertedID.(primitive.ObjectID), AccountType: newUser.AccountType, }, models.IError{}
}

func (ur *MongoUserRepository) Get(request interface{}) (IUser, models.IError) {
	ctx := context.TODO();
	userResponse:=IUser{}

	err := ur.repo.FindOne(ctx, &request).Decode(&userResponse);
	if (err != nil) {
		return IUser{}, models.IError{ Message: "unable to get this user" }
	}

	return userResponse, models.IError{};
}

func (ur *MongoUserRepository) GetMultiple(search IUser, option models.IOptions) (IMultipleUser, models.IError) {
	ctx := context.TODO();
	var users []IUser;
	option.Page = (option.Page - 1) * option.Limit

	userResponse, findErr := ur.repo.Find(ctx, search,
		&options.FindOptions{ Limit: &option.Limit, Skip: &option.Page },
	)
	if (findErr != nil) {
		return IMultipleUser{}, models.IError{ Message: "unable to get created deals" }
	}
	if decodeErr := userResponse.All(ctx, &users); decodeErr != nil {
		return IMultipleUser{}, models.IError{ Message: "unable to get created deal" }
	}

	responseCount, findErr := ur.repo.CountDocuments(ctx, search);
	if (findErr != nil) {
		return IMultipleUser{}, models.IError{ Message: "unable to count" }
	}

	multipleUserResponse := IMultipleUser{
		TotalUsers: responseCount,
		Users: users,
		HasNext: ((responseCount < ((option.Page + 1) * option.Limit))),
	}

	return multipleUserResponse, models.IError{};
}

func (ur *MongoUserRepository) Update(user interface{}, update interface{}) (IUser, models.IError) {
	ctx := context.TODO();

	var userResponse, err = ur.repo.UpdateOne(ctx, user, update);
	if (err != nil) {
		return IUser{}, models.IError{ Message: "`unable to get this user" }
	}

	return IUser{
		ID: userResponse.UpsertedID.(primitive.ObjectID),
	}, models.IError{};
}