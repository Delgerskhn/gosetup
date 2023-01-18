package persistence

import (
	"time"

	appgorm "github.com/delgerskhn/gosetup/pkg/domain/gorm"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	connect()
	configureDatabase()
	configureModels()
}

func connect() {
	dbConnectionString := viper.GetString("db_connection_string")
	var err error
	db, err = gorm.Open(sqlite.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

//instantiate gorm db
func GetDB() *gorm.DB {
	if db == nil {
		connect()
	}
	return db
}

func configureDatabase() {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func configureModels() error {
	return db.AutoMigrate(&appgorm.BookGorm{})
}
