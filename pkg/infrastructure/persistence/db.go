package persistence

import (
	"database/sql"
	"time"

	appgorm "github.com/delgerskhn/gosetup/pkg/domain/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

//instantiate gorm db
func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	_db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = _db

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}

	configureDatabase(sqlDB)
	configureModels(db)

	return db
}

func configureDatabase(sqlDB *sql.DB) {
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func configureModels(db *gorm.DB) error {
	return db.AutoMigrate(&appgorm.BookGorm{})
}
