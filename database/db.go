package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"todo-clean/common"
)

func ConnectDB() (client *gorm.DB, err error) {
	host := os.Getenv(common.PostgresHost)
	user := os.Getenv(common.PostgresUser)
	pw := os.Getenv(common.PostgresPassWord)
	port := os.Getenv(common.PostgresPort)
	dbName := os.Getenv(common.PostgresDB)
	ssl := os.Getenv(common.PostgresSSL)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbName, pw, ssl)

	client, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Postgres Client: error:")
		return nil, err
	}

	return client, err
}
