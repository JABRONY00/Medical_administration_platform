package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JABRONY00/medical_administration_platform/app/api/user-service/controllers"
)

var DB *pgxpool.Pool

func EmployeeRoutes(r *gin.Engine, db *pgxpool.Pool) {
	DB = db

	r.POST("/employees", controllers.CreateEmployee)
	r.GET("/employees", controllers.GetEmployees)
	r.GET("/employees/:id", controllers.GetEmployee)
	r.DELETE("/employees/:id", controllers.DeleteEmployee)
	r.PUT("/employees/:id", controllers.UpdateEmployee)
}
