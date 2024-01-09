package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/ydssx/morphix/pkg/errors"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewDB(dsn string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      NewGormLogger(zapcore.InfoLevel, zapcore.InfoLevel, time.Millisecond*200),
		PrepareStmt: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to mysql")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get mysql db")
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	
	return db, nil
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
