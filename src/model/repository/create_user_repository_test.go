package repository

import (
	"github.com/fabiokusaba/meu-primeiro-crud-go/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DATABASE", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain("test@email.com", "test123@", "test", 90)
		userDomain, err := repo.CreateUser(domain)

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), domain.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), domain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), domain.GetAge())
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain("test@email.com", "test123@", "test", 90)
		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
