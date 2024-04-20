package service

import (
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID service", zap.String("journey", "findUserByID"))
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail service", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail service", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
