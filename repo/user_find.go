package repo

import (
	"context"
	"gorm.io/gorm"
	"user-service/common"
	"user-service/model"
)

func (db *dbObj) FindUser(ctx context.Context, conditions map[string]interface{}) (*model.User, error) {
	var data model.User

	if err := db.db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewErrorResponse(404, "Not found", gorm.ErrRecordNotFound)
		}
		return nil, common.NewErrorResponse(500, "Database error", err)
	}
	return &data, nil
}
