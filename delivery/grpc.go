package delivery

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	__ "todo-clean/delivery/grpc"
	"todo-clean/domain"
)

func NewServerGrpc(grpcServer *grpc.Server, usecase domain.TodoUseCase) {
	todoServer := &newHandler{
		usecase: usecase,
	}
	__.RegisterTodoHandlerServer(grpcServer, todoServer)
	reflection.Register(grpcServer)
}

//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative todo-clean/delivery/grpc/todo.proto
