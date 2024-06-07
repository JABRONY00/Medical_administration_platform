package services

import (
	"reflect"

	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/models"
	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateClient(c *gin.Context, client *models.ClientWithPassword) error {
	//email verification
	client.ID = uuid.New().String()
	var err error
	client.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(client.Password), 10)
	if err != nil {
		return err
	}

	err = query.InsertClient(client)
	if err != nil {
		return err
	}

	return nil
}

func GetClient(clientID *string) (*models.ClientInfo, error) {
	client, err := query.GetClient(clientID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetClients(ids []string) ([]models.ClientInfo, error) {
	var (
		clients []models.ClientInfo
		err     error
	)

	if len(ids) == 0 {
		clients, err = query.GetAllClients()
	} else {
		clients, err = query.GetClientsByIDs(ids)
	}

	return clients, err
}

func DeleteClient(id *string) error {
	err := query.DeleteClient(id)
	return err
}

func UpdateClient(updClient *models.ClientInfo) error {
	oldClient, err := GetClient(&updClient.ID)
	if err != nil {
		return err
	}

	updClientReflection := reflect.ValueOf(&updClient).Elem()
	oldClientReflection := reflect.ValueOf(&oldClient).Elem()
	for i := 1; i < updClientReflection.NumField(); i++ {
		if updClientReflection.Field(i).IsZero() {
			updClientReflection.Field(i).Set(oldClientReflection.Field(i))
		}
	}

	err = query.UpdateClient(updClient)

	return err
}
