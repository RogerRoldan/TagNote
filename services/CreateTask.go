package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateTask(db *gorm.DB, task models.Task) {
	db.Create(&task)
	log.Println("Tarea creada exitosamente...")
}

func ReadTask(db *gorm.DB, task models.Task) {
	db.Find(&task)
	log.Println("Tarea encontrada exitosamente...")
}

func DeleteTask(db *gorm.DB, task models.Task) {
	db.Delete(&task)
	log.Println("Tarea eliminada exitosamente...")
}

func UpdateTask(db *gorm.DB, task models.Task) {
	db.Save(&task)
	log.Println("Tarea actualizada exitosamente...")
}
