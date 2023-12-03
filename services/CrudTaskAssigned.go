package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateTaskUser(db *gorm.DB, taskUser models.TaskAssigned) {
	db.Create(&taskUser)
	log.Println("Usuario de tarea creado exitosamente...")
}

func ReadTaskUser(db *gorm.DB, taskUser models.TaskAssigned) {
	db.Find(&taskUser)
	log.Println("Usuario de tarea encontrado exitosamente...")
}

func DeleteTaskUser(db *gorm.DB, taskUser models.TaskAssigned) {
	db.Delete(&taskUser)
	log.Println("Usuario de tarea eliminado exitosamente...")
}

func UpdateTaskUser(db *gorm.DB, taskUser models.TaskAssigned) {
	db.Save(&taskUser)
	log.Println("Usuario de tarea actualizado exitosamente...")
}
