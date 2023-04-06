package config

import "gorm.io/gorm"

type AppContext interface {
	GetDatabaseConnection() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (a appContext) GetDatabaseConnection() *gorm.DB {
	return a.db
}
