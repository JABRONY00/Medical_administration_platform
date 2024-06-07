package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/controllers"
)

func EmployeeRoutes(r *gin.Engine) {

	r.POST("/employees", controllers.CreateEmployee)
	r.GET("/employees", controllers.GetEmployees)
	r.GET("/employees/:id", controllers.GetEmployee)
	r.DELETE("/employees/:id", controllers.DeleteEmployee)
	r.PUT("/employees/:id", controllers.UpdateEmployee)
}
