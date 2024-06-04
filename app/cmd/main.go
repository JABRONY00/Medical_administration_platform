package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/JABRONY00/medical_administration_platform/app/api"
	"github.com/JABRONY00/medical_administration_platform/app/helpers"
	"github.com/JABRONY00/medical_administration_platform/app/initializers"
	"github.com/gin-gonic/gin"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")
var SERVER_HOST = helpers.GetEnv("SERVER_HOST")

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

func main() {
	router := gin.Default()
	initializers.ConnectDb()
	api.Routes(router)

	err := router.Run(fmt.Sprintf("%v:%v", SERVER_HOST, SERVER_PORT))

	if err != nil {
		log.Panicf("Failed to start server!: %v", err.Error())
	}
}
