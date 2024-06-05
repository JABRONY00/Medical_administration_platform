package employees

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	DB = db

	r.POST("/employees", CreateEmployee)
	r.GET("/employees", GetEmployees)
	r.GET("/employees/:id", GetEmployee)
	r.DELETE("/employees/:id", DeleteEmployee)
	r.PUT("/employees/:id", UpdateEmployee)
}
