package database

import (
	"ais.com/m/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost " +
		"user=kotyarich " +
		"password=1234 " +
		"dbname=postgres port=5432 TimeZone=Europe/Moscow"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Gun{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
