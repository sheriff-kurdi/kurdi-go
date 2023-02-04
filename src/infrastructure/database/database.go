package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB
var SqliteDB *gorm.DB

func Connect() {
	dsn := "host=10.5.0.2 user=postgres password=123456789 dbname=kurdi_go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
