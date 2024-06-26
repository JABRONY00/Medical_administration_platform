package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JABRONY00/medical_administration_platform/app/api/user-service/clients"
	"github.com/JABRONY00/medical_administration_platform/app/api/user-service/employees"
)

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	InitMiddlewares(r)

	employees.Routes(r, db)
	clients.Routes(r, db)
}
