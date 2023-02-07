package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

func (repo newRepo) Begin(ctx context.Context) (tx *gorm.DB, err error) {

	tx = repo.db.Begin()
	err = tx.Error

	if err != nil {
		return nil, fmt.Errorf("begin repository error: %w", err)
	}

	return tx, nil
}

func (repo newRepo) RollBack(db *gorm.DB) (err error) {

	err = db.Rollback().Error

	if err != nil {
		return err
	}

	return nil
}

func (repo newRepo) Commit() (tx *gorm.DB, err error) {

	tx = repo.db.Commit()
	err = tx.Error

	if err != nil {
		return nil, err
	}

	return tx, nil
}
