package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB
var SqliteDB *gorm.DB

func Connect() {

	database, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	PostgresDB = database

	SqliteDb, SqliteErr := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if SqliteErr != nil {
		panic(err)
	}
	SqliteDB = SqliteDb

}
