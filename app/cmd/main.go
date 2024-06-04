package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/JABRONY00/medical_administration_platform/app/api"
	"github.com/JABRONY00/medical_administration_platform/app/initializers"
	"github.com/gin-gonic/gin"
)

var SERVER_PORT = 4000

func init() {

	initializers.InitLogger()
	//Dbpool := initializers.DbConnection()
}

func main() {
	router := gin.Default()
	api.Routes(router)
	err := router.Run(fmt.Sprintf("Starting server on port: %v ...", SERVER_PORT))
	if err != nil {
		log.Panicf("Failed to start server!: %v", err.Error())
	}
}
