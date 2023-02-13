package repository

import (
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/lib/errorLib"
)

func (repo newRepo) Begin() (tx *gorm.DB, err error) {

	tx = repo.db.Begin()
	err = tx.Error

	if err != nil {
		return nil, errorLib.WrapError(common.ErrBeginTodo.Error(), err)
	}

	return tx, nil
}
