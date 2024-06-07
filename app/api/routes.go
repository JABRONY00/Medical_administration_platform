package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JABRONY00/medical_administration_platform/app/api/user-service/routes"
)

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	InitMiddlewares(r)

	routes.ClientRoutes(r, db)
	routes.EmployeeRoutes(r, db)
}
