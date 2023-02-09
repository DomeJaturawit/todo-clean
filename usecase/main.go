package usecase

import (
	"todo-clean/domain"
)

type newUseCase struct {
	repo domain.TodoRepositoryInterface
}

func NewUseCase(repo domain.TodoRepositoryInterface) domain.TodoUseCaseInterface {
	return &newUseCase{repo: repo}
}
