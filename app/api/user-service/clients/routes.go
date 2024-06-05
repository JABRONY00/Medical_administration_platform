package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	DB = db

	r.POST("/employees", CreateClient)
	r.GET("/employees", GetClients)
	r.GET("/employees/:id", GetClient)
	r.DELETE("/employees/:id", DeleteClient)
	r.PUT("/employees/:id", UpdateClient)
}
