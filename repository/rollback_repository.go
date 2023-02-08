package repository

import (
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/lib/error_lib"
)

func (repo newRepo) RollBack(db *gorm.DB) (err error) {

	err = db.Rollback().Error

	if err != nil {
		return error_lib.WrapError(common.ErrRollbackCreateTodo.Error(), err)
	}

	return nil
}
