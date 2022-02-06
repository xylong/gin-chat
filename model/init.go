package model

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// Database 连接数据库
func Database(dsn string) {
	var (
		db  *gorm.DB
		err error
	)

	loggerConfig := logger.Config{
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Error,
	}
	if gin.Mode() != "release" {
		loggerConfig.LogLevel = logger.Info
	}

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			loggerConfig,
		),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logrus.Info(err)
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	migrate()
}
