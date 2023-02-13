package repository

import (
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/lib/errorLib"
)

func (repo newRepo) Commit(db *gorm.DB) (err error) {

	err = db.Commit().Error

	if err != nil {
		return errorLib.WrapError(common.ErrCommitTodo.Error(), err)
	}

	return nil
}
