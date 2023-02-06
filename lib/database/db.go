package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var requiredEnvironmentVariablesForPostGres = []string{
	"POSTGRES_HOST",
	"POSTGRES_USER",
	"POSTGRES_PASSWORD",
	"POSTGRES_DB",
	"POSTGRES_PORT",
	"POSTGRES_SSL",
}

func ConnectDB() (client *gorm.DB, err error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	ssl := os.Getenv("POSTGRES_SSL")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbName, pw, ssl)

	client, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Postgres Client: error:")
		return nil, err
	}

	return client, err
}
