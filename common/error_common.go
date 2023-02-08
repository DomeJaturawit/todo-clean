package common

import "errors"

var (
	ErrDBCreateTodoRepo   = errors.New("failed to create todo repository in database")
	ErrDBCreateTodo       = errors.New("failed to create todo repository")
	ErrBeginCreateTodo    = errors.New("failed to begin create todo repository")
	ErrCommitCreateTodo   = errors.New("failed to commit create todo repository")
	ErrRollbackCreateTodo = errors.New("failed to roll back create todo repository")
	//UseCase
	ErrUseCaseCreateTodo = errors.New("failed to create todo usecase")

	//Dalivery
	ErrTitleEmpty       = errors.New("error input title is empty")
	ErrStatusEmpty      = errors.New("error input status is empty")
	ErrDescriptionEmpty = errors.New("error input description is empty")

	ErrFormat     = errors.New("invalid input format")
	ErrInternal   = errors.New("internal server error")
	ErrBadRequest = errors.New("bad request")
)
