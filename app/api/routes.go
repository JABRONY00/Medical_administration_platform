package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	usermicroservice "github.com/JABRONY00/medical_administration_platform/app/api/userMicroservice"
)

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	InitMiddlewares(r)
	usermicroservice.Routes(r, db)
}
