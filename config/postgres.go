package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetPostgersDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=job password=admin123")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
