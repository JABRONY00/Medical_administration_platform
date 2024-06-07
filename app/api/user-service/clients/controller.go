package clients

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	log "github.com/JABRONY00/medical_administration_platform/app/helpers/log"
)

func CreateClient(c *gin.Context) {
	var client ClientWithPassword
	err := c.ShouldBindJSON(&client)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "invalid request body")
		c.JSON(http.StatusBadRequest, "invalid request body")
		return
	}

	client.ID = uuid.New()

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

	var clients []ClientInfo
	for rows.Next() {
		var client ClientInfo
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
	var client ClientInfo
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

	var updclient ClientWithPassword
	err := c.ShouldBindJSON(&updclient)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "invalid request body")
		c.JSON(http.StatusBadRequest, "invalid request body")
		return
	}

	row := DB.QueryRow(context.Background(),
		"SELECT * FROM clients WHERE id = $1", id[0])
	var oldclient ClientInfo
	err = row.Scan(&oldclient.ID,
		&oldclient.FirstName,
		&oldclient.LastName,
		&oldclient.Phone,
		&oldclient.Email,
		&oldclient.PasswordHash,
		&oldclient.BirthDate,
		&oldclient.Gender,
		&oldclient.Address,
	)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	id = nil

	if updclient.Password != "" {
		updclient.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(updclient.Password), 10)
		if err != nil {
			log.HttpLog(c, log.Error, http.StatusInternalServerError, "failed to generate password hash")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	updclientReflection := reflect.ValueOf(&updclient.ClientInfo).Elem()
	oldclientReflection := reflect.ValueOf(&oldclient).Elem()
	for i := 1; i < updclientReflection.NumField(); i++ {
		if updclientReflection.Field(i).IsZero() {
			updclientReflection.Field(i).Set(oldclientReflection.Field(i))
		}
	}
	updclient.ID = oldclient.ID

	_, err = DB.Exec(context.Background(), "UPDATE clients SET first_name = $1, last_name = $2, phone = $3, email = $4, password_hash = $5, birth_date = $6, gender = $7, address = $8 WHERE id =$9",
		updclient.FirstName,
		updclient.LastName,
		updclient.Phone,
		updclient.Email,
		updclient.PasswordHash,
		updclient.BirthDate,
		updclient.Gender,
		updclient.Address,
		updclient.ID,
	)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	infoString := fmt.Sprintf("client with id=%s has been updated! Details: %s",
		oldclient.ID,
		"expected",
	)

	log.HttpLog(c, log.Info, http.StatusOK, infoString)
	c.AbortWithStatus(http.StatusOK)
}
