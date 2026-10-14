package db

import (
	"context"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Open(dsn string) (*gorm.DB, error) {
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, errors.New("postgres connection unavailable")
	}
	sql, err := d.DB()
	if err != nil {
		return nil, errors.New("postgres initialization failed")
	}
	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(20)
	sql.SetConnMaxLifetime(30 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = sql.PingContext(ctx); err != nil {
		_ = sql.Close()
		return nil, errors.New("postgres ping failed")
	}
	return d, nil
}
