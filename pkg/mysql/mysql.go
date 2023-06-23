package mysql

import (
	"time"

	"github.com/axiaoxin-com/logging"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(dsn string) *gorm.DB {
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logging.NewGormLogger(zapcore.InfoLevel, zapcore.InfoLevel, time.Second*2)})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return db
}
