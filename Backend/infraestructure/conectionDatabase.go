package infraestructure

import (
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
	var stringConnection = "host=localhost user=postgres password=root dbname=workhub port=5432 sslmode=disable TimeZone=America/Bogota"
	//var stringConnection = "host=localhost user=postgres password=postgres dbname=workhub port=5433 sslmode=disable TimeZone=America/Bogota"
>>>>>>> 7b381408b12f3a6219b107f4523d3cae726b20f3
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Conectado exitosamente a la DB...")

	return db
}
