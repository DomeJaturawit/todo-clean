package usecase

import (
	"context"
	"todo-clean/domain"
)

func (n *newUseCase) UpdateTodoUseCase(ctx context.Context, queryEntity domain.UpdateTodoQueryEntity, entity *domain.UpdateTodoEntity) (result *domain.UpdateTodoEntity, err error) {
	//
	//response, err := n.repo.GetTodoRepository(ctx, &queryEntity.ID)
	//if len(response) == 0 {
	//	return nil, errorLib.WrapError(common.ErrDataNotFound.Error(), err)
	//}
	//
	//if entity.Title == "" {
	//	entity.Title = response[0].Title
	//}
	//if entity.Status == "" {
	//	entity.Status = response[0].Status
	//}
	//if entity.Description == "" {
	//	entity.Description = response[0].Description
	//}
	//
	//dbTx, err := n.repo.Begin()
	//if err != nil {
	//	return nil, errorLib.WrapError(common.ErrBeginTodo.Error(), err)
	//}
	//result, err = n.repo.UpdateTodoRepository(ctx, dbTx, queryEntity, domain.UpdateTodoEntity{
	//	Title:       entity.Title,
	//	Description: entity.Description,
	//	Status:      entity.Status,
	//	UpdatedAt:   time.Now(),
	//})
	//if err != nil {
	//	return nil, errorLib.WrapError(common.ErrUseCaseUpdateTodo.Error(), err)
	//}
	//err = n.repo.Commit(dbTx)
	//if err != nil {
	//	if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {
	//		return nil, errorLib.WrapError(common.ErrRollbackTodo.Error(), rollbackErr)
	//	}
	//	return nil, errorLib.WrapError(common.ErrCommitTodo.Error(), err)
	//}

	return result, err
}
