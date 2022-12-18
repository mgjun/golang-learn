package dao

import (
	"go-admin/internal/app/config"
	"gorm.io/gorm"
	"strings"
)

func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnnoDB")
	}
	return db.AutoMigrate()
}
