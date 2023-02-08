package repository

import (
	"context"
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/lib/error_lib"
)

func (repo newRepo) Begin(ctx context.Context) (tx *gorm.DB, err error) {

	tx = repo.db.Begin()
	err = tx.Error

	if err != nil {
		return nil, error_lib.WrapError(common.ErrBeginCreateTodo.Error(), err)
	}

	return tx, nil
}
