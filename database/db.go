package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todo-clean/repository/model"
)

func ConnectDB() (client *gorm.DB, err error) {
	//TODO: Use ENV
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", "localhost", "5433", "pg", "crud", "pass", "disable")
	client, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	err = MigrateDB(client)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func MigrateDB(dbConn *gorm.DB) (err error) {
	err = dbConn.AutoMigrate(
		&model.TbTodoRepositoryCreateModel{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
