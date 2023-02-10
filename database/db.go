package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todo-clean/repository/model"
)

func ConnectDB() (client *gorm.DB, err error) {
	//host := os.Getenv(common.PostgresHost)
	//user := os.Getenv(common.PostgresUser)
	//pw := os.Getenv(common.PostgresPassWord)
	//port := os.Getenv(common.PostgresPort)
	//dbName := os.Getenv(common.PostgresDB)
	//ssl := os.Getenv(common.PostgresSSL)

	//- POSTGRES_HOST=localhost
	//- POSTGRES_USER=pg
	//- POSTGRES_PORT=54321
	//- POSTGRES_PASSWORD=pass
	//- POSTGRES_DB=crud
	//- POSTGRES_SSL=disable
	//log.Println("host ==>>", host)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", "localhost", "5433", "pg", "crud", "pass", "disable")
	client, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("err =>>", err)
		return nil, err
	}

	err = MigrateDB(client)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func MigrateDB(dbConn *gorm.DB) (err error) {
	err = dbConn.AutoMigrate(
		&model.TbTodoRepositoryModel{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
