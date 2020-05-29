package wxdb

import (
	"github.com/jinzhu/gorm"
)

var wxDB *gorm.DB

func DB() *gorm.DB {
	return wxDB
}

func SetDB(db *gorm.DB) {
	wxDB = db
}
