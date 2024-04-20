package controller

import (
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller", zap.String("journey", "deleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call DeleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully", zap.String("userId", userId), zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)

}
