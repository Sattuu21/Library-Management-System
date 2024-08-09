package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=postgres password=satyam12 dbname=postgres port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	logger.Default.LogMode(logger.Info)
	if err!= nil{
		panic(err)	
	}
	db =d
}

func GetDB() *gorm.DB{
	return db
}
