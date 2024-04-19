package repository

import (
	"context"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/model/repository/entity/converter"
	"go.uber.org/zap"
	"os"

	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/model"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	logger.Info("CreateUser repository executed successfully", zap.String("userId", userDomain.GetID()), zap.String("journey", "createUser"))

	return userDomain, nil
}
