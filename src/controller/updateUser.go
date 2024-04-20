package controller

import (
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("journey", "updateUser"))

	var userRequest request.UserUpdateRequest

	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "updateUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service", err, zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully", zap.String("userId", userId), zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
