package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"todo-clean/database"
	"todo-clean/delivery"
	"todo-clean/repository"
	"todo-clean/usecase"
)

func main() {

	dbConn, err := database.ConnectDB()

	if err != nil {
		log.Panicln("failed to connect database", err)
	} else {
		fmt.Println("Connect ok", dbConn)
	}
	//engine := gin.New()

	repo := repository.NewRepository(dbConn)
	use := usecase.NewUseCase(repo)

	lis, err := net.Listen("tcp", "localhost:9000")

	server := grpc.NewServer()
	delivery.NewServerGrpc(server, use)

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("started")

	//err = engine.Run("localhost:8080")

	if err != nil {
		log.Fatalln("failed to run", err)
	}
}
