package repository

import (
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/lib/error_lib"
)

func (repo newRepo) Commit(db *gorm.DB) (err error) {

	err = db.Commit().Error

	if err != nil {
		return error_lib.WrapError(common.ErrCommitCreateTodo.Error(), err)
	}

	return nil
}
