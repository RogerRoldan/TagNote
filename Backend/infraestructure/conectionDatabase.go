package infraestructure

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// function return connection database
func GetConnection() *gorm.DB {
	params := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, DB, PORT)
	var stringConnection = params

	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Conectado exitosamente a la DB...")

	return db
}
