package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/controllers"
)

func ClientRoutes(r *gin.Engine) {

	r.POST("/clients", controllers.CreateClient)
	r.GET("/clients", controllers.GetClients)
	r.GET("/clients/:clientId", controllers.GetClient)
	r.DELETE("/clients/:clientId", controllers.DeleteClient)
	r.PUT("/clients/:client_Id", controllers.UpdateClient)
}
