package persistence

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

//instantiate gorm db
func GetDB() *gorm.DB {
	_db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db = _db
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
