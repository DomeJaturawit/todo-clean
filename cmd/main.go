package main

import (
	"fmt"
	"log"
	"todo-clean/database"
	"todo-clean/delivery"
	"todo-clean/repository"
	"todo-clean/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConn, err := database.ConnectDB()
	//
	if err != nil {
		log.Panicln("failed to connect database", err)
	} else {
		fmt.Println("Connect ok", dbConn)
	}
	engine := gin.New()

	repo := repository.NewRepository(dbConn)
	use := usecase.NewUseCase(repo)
	//
	delivery.NewHandler(engine, use)
	//
	err = engine.Run("localhost:8080")
	if err != nil {
		log.Fatalln("failed to run", err)
	}
}
