package infraestructure

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// function return connection database
func GetConnection() *gorm.DB {
<<<<<<< HEAD

	//var stringConnection = "host=localhost user=postgres password=root dbname=workhub port=5433 sslmode=disable TimeZone=America/Bogota"
	var stringConnection = "host=localhost user=postgres password=postgres dbname=workhub port=5432 sslmode=disable TimeZone=America/Bogota"
=======
	params := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, DB, PORT)
	var stringConnection = params
>>>>>>> a1bfd930df910eca29298becc43b3bb79b65f6b0
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Conectado exitosamente a la DB...")

	return db
}
