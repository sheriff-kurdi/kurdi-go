package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123456789 dbname=kurdi_go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	postgresDb, postgresErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if postgresErr != nil {
		panic(postgresErr)
	}
	PostgresDB = postgresDb

}
