package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	DB = db

	r.POST("/clients", CreateClient)
	r.GET("/clients", GetClients)
	r.GET("/clients/:id", GetClient)
	r.DELETE("/clients/:id", DeleteClient)
	r.PUT("/clients/:id", UpdateClient)
}
