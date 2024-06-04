package initializers

import (
	"context"
	"fmt"

	"github.com/JABRONY00/medical_administration_platform/app/helpers"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	DB_PORT     = helpers.GetEnv("DB_PORT")
	DB_HOST     = helpers.GetEnv("DB_HOST")
	DB_NAME     = helpers.GetEnv("DB_NAME")
	DB_USER     = helpers.GetEnv("DB_USER")
	DB_PASSWORD = helpers.GetEnv("DB_PASSWORD")
)

func ConnectDb() *pgxpool.Pool {
	postgreURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	dbpool, err := pgxpool.New(context.Background(), postgreURL)
	if err != nil {
		log.Panicf("DB Connection failed: %v", err.Error())
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Panicf("DB Ping failed: %v", err.Error())
	}

	log.Infof("DB connected successfully on port: %v", DB_PORT)
	return dbpool
}
