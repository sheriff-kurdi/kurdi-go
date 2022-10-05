package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB
var SqliteDB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123456789 dbname=kurdi_go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	postgresDb, postgresErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if postgresErr != nil {
		panic(postgresErr)
	}
	PostgresDB = postgresDb

	sqliteDb, sqliteErr := gorm.Open(sqlite.Open("kurdi_go.db"), &gorm.Config{})
	if sqliteErr != nil {
		panic(sqliteErr)
	}
	SqliteDB = sqliteDb

}
