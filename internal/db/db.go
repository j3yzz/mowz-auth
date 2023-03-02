package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	MAX_IDLE_CONNECTIONS   = 10
	MAX_OPEN_CONNECTIONS   = 100
	MAX_CONNECTION_TIMEOUT = 10 * time.Second
)

func New(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("db new client error: %w", err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	sqlDB.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
	sqlDB.SetConnMaxLifetime(MAX_CONNECTION_TIMEOUT)

	return db, nil
}
