package models

import (
	"github.com/Gre-Z/common/jtime"
	"github.com/ydssx/morphix/pkg/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint            `json:"id" gorm:"primary_key"`
	CreatedAt jtime.JsonTime  `json:"created_at"`
	UpdatedAt jtime.JsonTime  `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

type DB struct {
	db *gorm.DB
}

func getDB(tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 {
		return tx[0]
	}
	return mysql.DB
}
