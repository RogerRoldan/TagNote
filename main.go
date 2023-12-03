package main

import (
	"fmt"
	_ "fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Cod        uint8  `gorm:"unique; primaryKey"`
	Name       string `gorm:"not null"`
	SecondName string
	LastName   string `gorm:"not null; default:'NN'"`
	Visited_at time.Time
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm_db port=5433 sslmode=disable TimeZone=America/Bogota"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Conectado exitosamente a la DB...")

	// Migrate the schema
	db.AutoMigrate(&Student{})

	//db.Delete(&Student{}).Error // gorm.ErrMissingWhereClause
	db.Where("1 = 1").Delete(&Student{})
	//DELETE FROM `students` WHERE 1=1

	db.Exec("DELETE FROM students")
	//DELETE FROM students

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Student{})
	//DELETE FROM students

	fmt.Println(db)

	// Create
	db.Create(&Student{Cod: 190, Name: "Pepito"})
	db.Create(&Student{Cod: 191, Name: "Anita", Visited_at: time.Now()})
	db.Create(&Student{Cod: 192, Name: "Paquita", Visited_at: time.Now()})

	// Read
	var student, student2, student3 Student
	db.First(&student, 1)               // find product with integer primary key
	db.First(&student2, "cod = ?", 192) // find product with code 192
	db.First(&student3, "cod = ?", 191) // find product with code 191
	fmt.Println(student)
	fmt.Println(student2)
	fmt.Println(student3)

	// Update - update student's Name to Marcos
	db.Model(&student).Update("name", "Marcos")
	// Update - update multiple fields
	db.Model(&student2).Updates(Student{Cod: 200, Name: "Margarita", LastName: "Rosa"}) // non-zero fields
	db.Model(&student3).Updates(map[string]interface{}{"Cod": 202, "Name": "Francisca", "LastName": "Rosas"})

	// Delete - delete student
	db.Delete(&student)
	db.Delete(&student2, 3)
	db.Delete(&Student{}, 1)
	students := []Student{student, student2, student3}
	db.Delete(&students, []int{1, 3})
	db.Delete(&Student{}, []int{1, 3})

}
