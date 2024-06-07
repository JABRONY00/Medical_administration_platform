package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/models"
	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/services"
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
	err = services.CreateClient(c, &client)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.HttpLog(c, log.Info, http.StatusOK, fmt.Sprintf("new client created, id=%s", client.ID))
	c.AbortWithStatus(http.StatusOK)
}

func GetClients(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	ids := queryParams["id"]

	clients, err := services.GetClients(ids)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, clients)
}

func GetClient(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	ids, ok := queryParams["id"]
	if !ok {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "missing id")
		c.JSON(http.StatusBadRequest, "missing id")
		return
	}

	client, err := services.GetClient(&ids[0])
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
	err := services.DeleteClient(&id[0])
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

	var updClient models.ClientInfo
	err := c.ShouldBindJSON(&updClient)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusBadRequest, "invalid request body")
		c.JSON(http.StatusBadRequest, "invalid request body")
		return
	}
	updClient.ID = id[0]

	err = services.UpdateClient(&updClient)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.HttpLog(c, log.Info, http.StatusOK, fmt.Sprintf("Client with id=%s has been updated!", updClient.ID))
	c.AbortWithStatus(http.StatusOK)
}
