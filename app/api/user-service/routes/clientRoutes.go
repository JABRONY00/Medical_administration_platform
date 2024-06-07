package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JABRONY00/medical_administration_platform/app/api/user-service/controllers"
)

var DB *pgxpool.Pool

func ClientRoutes(r *gin.Engine, db *pgxpool.Pool) {
	DB = db

	r.POST("/clients", controllers.CreateClient)
	r.GET("/clients", controllers.GetClients)
	r.GET("/clients/:clientId", controllers.GetClient)
	r.DELETE("/clients/:clientId", controllers.DeleteClient)
	r.PUT("/clients/:client_Id", controllers.UpdateClient)
}
