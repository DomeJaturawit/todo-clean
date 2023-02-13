package usecase

import (
	"todo-clean/domain"
)

type newUseCase struct {
	repo domain.TodoRepository
}

func NewUseCase(repo domain.TodoRepository) domain.TodoUseCase {
	return &newUseCase{repo: repo}
}
