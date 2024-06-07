package controllers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/models"
	log "github.com/JABRONY00/medical_administration_platform/app/helpers/log"
)

func CreateClient(c *gin.Context) {
	var client models.ClientWithPassword
	err := c.ShouldBindJSON(&client)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "invalid request body")
		c.JSON(http.StatusBadRequest, "invalid request body")
		return
	}

	switch {
	case client.FirstName == "":
		fallthrough
	case client.LastName == "":
		fallthrough
	case !(client.Gender == "M" || client.Gender == "F"):
		fallthrough
	case client.Email == "":
		{
			log.HttpLog(c, log.Error, http.StatusBadRequest, "invalid request body")
			c.JSON(http.StatusBadRequest, "invalid request body")
			return
		}
	}

	client.ID = uuid.New().String()

	client.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(client.Password), 10)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, "failed to generate password hash")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = DB.Exec(context.Background(),
		"INSERT INTO clients VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		client.ID,
		client.FirstName,
		client.LastName,
		client.Phone,
		client.Email,
		client.PasswordHash,
		client.BirthDate,
		client.Gender,
		client.Address,
	)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.HttpLog(c, log.Info, http.StatusOK, fmt.Sprintf("new client created, id=%s", client.ID.String()))
	c.AbortWithStatus(http.StatusOK)
}

func GetClients(c *gin.Context) {
	fmt.Println("get clients endpoint")
	rows, err := DB.Query(context.Background(),
		"SELECT * FROM clients")
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clients []models.ClientInfo
	for rows.Next() {
		var client models.ClientInfo
		err := rows.Scan(&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Phone,
			&client.Email,
			&client.PasswordHash,
			&client.BirthDate,
			&client.Gender,
			&client.Address,
		)
		if err != nil {
			log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		clients = append(clients, client)
	}

	c.JSON(http.StatusOK, clients)
}

func GetClient(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	id, ok := queryParams["id"]
	if !ok {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "missing id")
		c.JSON(http.StatusBadRequest, "missing id")
		return
	}

	row := DB.QueryRow(context.Background(),
		"SELECT * FROM clients WHERE id = $1", id[0])
	var client models.ClientInfo
	err := row.Scan(&client.ID,
		&client.FirstName,
		&client.LastName,
		&client.Phone,
		&client.Email,
		&client.PasswordHash,
		&client.BirthDate,
		&client.Gender,
		&client.Address,
	)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, client)
}

func DeleteClient(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	id, ok := queryParams["id"]
	if !ok {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "missing id")
		c.JSON(http.StatusBadRequest, "missing id")
		return
	}

	_, err := DB.Exec(context.Background(),
		"DELETE FROM clients WHERE id=$1", id[0])
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.HttpLog(c, log.Info, http.StatusOK, fmt.Sprintf("client with id=%s has been deleted", id[0]))
	c.AbortWithStatus(http.StatusOK)
}

func UpdateClient(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	id, ok := queryParams["id"]
	if !ok {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "missing id")
		c.JSON(http.StatusBadRequest, "missing id")
		return
	}

	var updClient models.ClientWithPassword
	err := c.ShouldBindJSON(&updClient)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "invalid request body")
		c.JSON(http.StatusBadRequest, "invalid request body")
		return
	}

	row := DB.QueryRow(context.Background(),
		"SELECT * FROM clients WHERE id = $1", id[0])
	var oldClient models.ClientInfo
	err = row.Scan(&oldClient.ID,
		&oldClient.FirstName,
		&oldClient.LastName,
		&oldClient.Phone,
		&oldClient.Email,
		&oldClient.PasswordHash,
		&oldClient.BirthDate,
		&oldClient.Gender,
		&oldClient.Address,
	)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if updClient.Password != "" {
		updClient.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(updClient.Password), 10)
		if err != nil {
			log.HttpLog(c, log.Error, http.StatusInternalServerError, "failed to generate password hash")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	updClientReflection := reflect.ValueOf(&updClient.ClientInfo).Elem()
	oldClientReflection := reflect.ValueOf(&oldClient).Elem()
	for i := 1; i < updClientReflection.NumField(); i++ {
		if updClientReflection.Field(i).IsZero() {
			updClientReflection.Field(i).Set(oldClientReflection.Field(i))
		}
	}
	updClient.ID = oldClient.ID

	_, err = DB.Exec(context.Background(), "UPDATE clients SET first_name = $1, last_name = $2, phone = $3, email = $4, password_hash = $5, birth_date = $6, gender = $7, address = $8 WHERE id =$9",
		updClient.FirstName,
		updClient.LastName,
		updClient.Phone,
		updClient.Email,
		updClient.PasswordHash,
		updClient.BirthDate,
		updClient.Gender,
		updClient.Address,
		updClient.ID,
	)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	infoString := fmt.Sprintf("client with id=%s has been updated! Details: %s",
		oldClient.ID,
		"expected",
	)

	log.HttpLog(c, log.Info, http.StatusOK, infoString)
	c.AbortWithStatus(http.StatusOK)
}
