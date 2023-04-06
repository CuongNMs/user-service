package repo

import "gorm.io/gorm"

type dbObj struct {
	db *gorm.DB
}

func NewDbObj(db *gorm.DB) *dbObj {
	return &dbObj{db: db}
}
