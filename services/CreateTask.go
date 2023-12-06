package services

import (
	"log"

	"github.com/roger/workhub/models"
	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB, task models.Task) (string, bool) {
	taskCreate := db.Create(&task)
	if taskCreate.Error != nil {
		log.Println("Error al crear la tarea")
		return "Error al crear la tarea", false
	}
	log.Println("Tarea creada exitosamente...")
	return "Tarea creada exitosamente...", true
}

func GetTasks(db *gorm.DB) ([]models.Task, bool) {
	var tasks []models.Task
	tasksList := db.Find(&tasks)
	if tasksList.Error != nil {
		log.Println("Error al obtener las tareas")
		return []models.Task{}, false
	}

	log.Println("Tareas encontradas exitosamente...")
	return tasks, true
}

func GetTaskId(db *gorm.DB, id int) (models.Task, bool) {
	var task models.Task
	response := db.First(&task, id)
	if response.Error != nil {
		log.Println("Tarea no encontrada...")
		return models.Task{}, false
	}

	log.Println("Tarea encontrada exitosamente...")
	return task, true
}

func DeleteTask(db *gorm.DB, id int) (string, bool) {
	var task models.Task
	var aux models.Task = models.Task{}
	db.First(&task, id)
	if task.ID == (aux.ID) {
		log.Println("Tarea no encontrada...")
		return "Tarea no encontrada...", false
	}

	taskDelete := db.Delete(&task)
	if taskDelete.Error != nil {
		log.Println("Error al eliminar la tarea")
		return "Error al eliminar la tarea", false
	}
	log.Println("Tarea eliminada exitosamente...")
	return "Tarea eliminada exitosamente...", true
}

func UpdateTask(db *gorm.DB, task models.Task) (string, bool) {
	taskM := db.Save(&task)
	if taskM.Error != nil {
		log.Println("Error al actualizar la tarea")
		return "Error al actualizar la tarea", false
	}
	log.Println("Tarea actualizada exitosamente...")
	return "Tarea actualizada exitosamente...", true
}
