package common

import "errors"

var (
	ErrDBCreateTodoRepo = errors.New("failed to create todo repository in database")
	ErrDBCreateTodo     = errors.New("failed to create todo repository")
	ErrDBUpdateTodo     = errors.New("failed to update todo repository")
	ErrBeginTodo        = errors.New("failed to begin todo repository")
	ErrCommitTodo       = errors.New("failed to commit create todo repository")
	ErrRollbackTodo     = errors.New("failed to roll back create todo repository")

	ErrDBGetTodo = errors.New("failed to get todo database repository")

	ErrGetAllTodo = errors.New("failed to get all todo")
	ErrGetTodo    = errors.New("failed to get todo")

	ErrDataNotFound = errors.New("data not found")
	//UseCase
	ErrUseCaseCreateTodo = errors.New("failed to create todo usecase")

	//Dalivery
	ErrTitleEmpty       = errors.New("error input title is empty")
	ErrStatusEmpty      = errors.New("error input status is empty")
	ErrDescriptionEmpty = errors.New("error input description is empty")

	ErrFormat     = errors.New("invalid input format")
	ErrInternal   = errors.New("internal server error")
	ErrBadRequest = errors.New("bad request")

	ErrCopierCopy = errors.New("failed to copier copy ")

	ErrSaveDB = errors.New("failed to save update to database")
)
