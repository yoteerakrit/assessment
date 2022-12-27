package repo

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		"0.0.0.0",
		"root",
		"expense",
		"temp",
		5455,
	)

	//docker run --name assessment -e POSTGRES_USER=root -e POSTGRES_PASSWORD=expense -e POSTGRES_DB=temp -p 5455:5432 -d postgres

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database : %s", err.Error())
	}

	return db
}
