package common

import "errors"

const (
//ErrRepository
//ErrDBCreateTodo       = ("failed to create todo repository")
//ErrBeginCreateTodo    = ("failed to begin create todo repository")
//ErrCommitCreateTodo   = ("failed to commit create todo repository")
//ErrRollbackCreateTodo = ("failed to roll back create todo repository")

// ErrUseCase
// ErrUseCaseCreateTodo = ("failed to create todo usecase")
)

var (
	ErrDBCreateTodoRepo   = errors.New("failed to create todo repository in database")
	ErrDBCreateTodo       = errors.New("failed to create todo repository")
	ErrBeginCreateTodo    = errors.New("failed to begin create todo repository")
	ErrCommitCreateTodo   = errors.New("failed to commit create todo repository")
	ErrRollbackCreateTodo = errors.New("failed to roll back create todo repository")

	ErrUseCaseCreateTodo = errors.New("failed to create todo usecase")
)
