package repo

import (
	"context"
	"user-service/common"
	"user-service/model"
)

func (db *dbObj) CreateUser(data *model.User, ctx context.Context) error {
	tx := db.db.Begin()
	if err := tx.Table(data.TableName()).Create(data).Error; err != nil {
		tx.Rollback()
		return common.NewErrorResponse(500, "Error when create", err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return common.NewErrorResponse(500, "Error when commit", err)
	}
	return nil
}
