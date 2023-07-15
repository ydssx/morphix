package mysql

import (
	"database/sql"
	"time"

	"github.com/axiaoxin-com/logging"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	DB = db
	return db
}

func Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return DB.Transaction(fc, opts...)
}
