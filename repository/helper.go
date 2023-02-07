package repository

import (
	"context"
	"gorm.io/gorm"
)

func (repo NewRepo) Begin(ctx context.Context) (tx *gorm.DB, err error) {

	tx = repo.db.Begin()
	err = tx.Error

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (repo NewRepo) RollBack() (tx *gorm.DB, err error) {

	tx = repo.db.Rollback()
	err = tx.Error

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (repo NewRepo) Commit() (tx *gorm.DB, err error) {

	tx = repo.db.Rollback()
	err = tx.Error

	if err != nil {
		return nil, err
	}

	return tx, nil
}
