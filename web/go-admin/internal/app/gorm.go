package app

import (
	"errors"
	"go-admin/internal/app/config"
	"go-admin/internal/app/dao"
	"go-admin/pkg/gormx"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func InitGormDB() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, err := NewGormDB()
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	if cfg.EnableAutoMigrate {
		err = dao.AutoMigrate(db)
		if err != nil {
			return nil, cleanFunc, err
		}
	}
	return db, cleanFunc, nil
}

func NewGormDB() (*gorm.DB, error) {
	cfg := config.C
	var dsn string
	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		_ = os.Mkdir(filepath.Dir(dsn), 0777)
	case "postgres":
		dsn = cfg.Postgres.DSN()
	default:
		return nil, errors.New("unknown db")
	}

	return gormx.New(&gormx.Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifeTime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
		TablePrefix:  cfg.Gorm.TablePrefix,
	})
}
