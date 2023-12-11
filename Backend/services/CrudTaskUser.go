package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateTaskUser(db *gorm.DB, taskUser *models.TaskUser) (string, bool) {

	//validate if tuple exists
	taskUserExists := db.Where("task_id = ? AND user_id = ?", taskUser.TaskID, taskUser.UserID).Find(&models.TaskUser{})
	if taskUserExists.RowsAffected > 0 {
		log.Println("Tarea ya asignada a usuario")
		return "Tarea ya asignada a usuario", false
	}

	taskUserCreated := db.Create(taskUser)
	if taskUserCreated.Error != nil {
		log.Println("Error al asignar tarea a usuario: ")
		return "Error al asignar tarea a usuario", false
	}
	log.Println("Tarea asignada a usuario correctamente")
	return "Tarea asignada a usuario correctamente", true
}

func GetTaskUser(db *gorm.DB) ([]models.TaskUser, bool) {
	var taskUsers []models.TaskUser
	taskUserList := db.Find(&models.TaskUser{}).Preload("User").Preload("Task").Find(&taskUsers)
	if taskUserList.Error != nil {
		log.Println("Error al obtener las tareas")
		return []models.TaskUser{}, false
	}

	log.Println("Tareas encontradas exitosamente...")
	return taskUsers, true
}

func GetTaskUserId(db *gorm.DB, id int) (models.TaskUser, bool) {
	var taskUser models.TaskUser
	response := db.First(&taskUser, id)
	if response.Error != nil {
		log.Println("Tarea no encontrada...")
		return models.TaskUser{}, false
	}

	log.Println("Tarea encontrada exitosamente...")
	return taskUser, true
}

func DeleteTaskUser(db *gorm.DB, id int) (string, bool) {
	var taskUser models.TaskUser
	var aux models.TaskUser = models.TaskUser{}
	db.First(&taskUser, id)
	if taskUser.ID == (aux.ID) {
		log.Println("Tarea no encontrada...")
		return "Tarea no encontrada...", false
	}

	taskUserDelete := db.Delete(&taskUser)
	if taskUserDelete.Error != nil {
		log.Println("Error al eliminar la tarea")
		return "Error al eliminar la tarea", false
	}
	log.Println("Tarea eliminada exitosamente...")
	return "Tarea eliminada exitosamente...", true
}

func GetTaskUserByUserId(db *gorm.DB, id int) ([]models.TaskUser, bool) {
	var taskUsers []models.TaskUser
	taskUserList := db.Where("user_id = ?", id).Find(&models.TaskUser{}).Preload("Task").Find(&taskUsers)
	if taskUserList.Error != nil {
		log.Println("Error al obtener las tareas")
		return []models.TaskUser{}, false
	}

	log.Println("Tareas encontradas exitosamente...")
	return taskUsers, true
}

func GetTaskUserByTaskId(db *gorm.DB, id int) ([]models.TaskUser, bool) {
	var taskUsers []models.TaskUser
	taskUserList := db.Where("task_id = ?", id).Find(&models.TaskUser{}).Preload("User").Find(&taskUsers)
	if taskUserList.Error != nil {
		log.Println("Error al obtener las tareas")
		return []models.TaskUser{}, false
	}

	log.Println("Tareas encontradas exitosamente...")
	return taskUsers, true
}

func UpdateTaskUser(db *gorm.DB, taskUser *models.TaskUser) (string, bool) {
	taskUserUpdate := db.Save(taskUser)
	if taskUserUpdate.Error != nil {
		log.Println("Error al actualizar la tarea")
		return "Error al actualizar la tarea", false
	}
	log.Println("Tarea actualizada exitosamente...")
	return "Tarea actualizada exitosamente...", true
}
