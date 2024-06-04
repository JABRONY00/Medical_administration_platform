package initializers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	DB_PORT     = "5000"
	DB_HOST     = "localhost"
	DB_NAME     = "j"
	DB_OWNER    = "j"
	DB_PASSWORD = "j"
)

func DbConnection() *pgxpool.Pool {
	postgreURL := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_OWNER, DB_PASSWORD, DB_NAME)
	dbpool, err := pgxpool.New(context.Background(), postgreURL)
	if err != nil {
		log.Panicf("DB Connection failed: %v", err.Error())
	}
	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Panicf("DB Ping failed: %v", err.Error())
	}
	log.Info("DB connected successfully")
	return dbpool
}
