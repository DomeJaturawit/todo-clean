package repository

import (
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/lib/errorLib"
)

func (repo newRepo) RollBack(db *gorm.DB) (err error) {

	err = db.Rollback().Error

	if err != nil {
		return errorLib.WrapError(common.ErrRollbackTodo.Error(), err)
	}

	return nil
}
