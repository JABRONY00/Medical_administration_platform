package usermicroservice

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/query"
	"github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice/routes"
)

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	query.TransferDB(db)
	routes.ClientRoutes(r)
	routes.EmployeeRoutes(r)
}
