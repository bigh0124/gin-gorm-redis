package config

import (
	"fmt"
	"log"
	"time"

	"github.com/bigh0124/gin-gorm-redis/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDB *gorm.DB

func InitDB() error {
	dsn := AppConfig.Database.Dsn
	if dsn == "" {
		return fmt.Errorf("dsn must be provided")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect db: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get db instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to test connect db: %w", err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("failed to migration: %w", err)
	}

	err = db.AutoMigrate(&model.Article{})
	if err != nil {
		return fmt.Errorf("failed to migration: %w", err)
	}

	gormDB = db
	log.Println("db connect successful")
	return nil
}

func GetDB() *gorm.DB {
	return gormDB
}
