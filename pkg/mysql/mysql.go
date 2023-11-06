package mysql

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewDB(dsn string) *gorm.DB {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      NewGormLogger(zapcore.InfoLevel, zapcore.InfoLevel, time.Millisecond*200),
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	return db
}

func Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return db.Transaction(fc, opts...)
}

func GlobalDB() *gorm.DB {
	return db
}

type contextTxKey struct{}

func NewContextWithDB(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextTxKey{}, db)
}

func DBFromContext(ctx context.Context) *gorm.DB {
	return ctx.Value(contextTxKey{}).(*gorm.DB)
}
