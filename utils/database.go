package utils

import (
	"os"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func ConnectToDatabaseMysql() (*gorm.DB, error) {
	file := os.Getenv("DB_SQLITE_FILE")
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})

	return db, err
}

func CloseConnectToDatabaseMysql(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
