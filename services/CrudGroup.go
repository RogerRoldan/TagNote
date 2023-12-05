package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateGroup(db *gorm.DB, group models.Group) (string, bool) {
	groupCreate := db.Create(&group)
	if groupCreate.Error != nil {
		log.Println("Error al crear el grupo")
		return "Error al crear el grupo", false
	}
	log.Println("Grupo creado exitosamente...")
	return "Grupo creado exitosamente...", true
}

func GetGroups(db *gorm.DB) ([]models.Group, bool) {
	var groups []models.Group
	groupsList := db.Find(&groups)
	if groupsList.Error != nil {
		log.Println("Error al obtener los grupos")
		return []models.Group{}, false
	}

	log.Println("Grupos encontrados exitosamente...")
	return groups, true
}

func GetGroupId(db *gorm.DB, id int) (models.Group, bool) {
	var group models.Group
	response := db.First(&group, id)
	if response.Error != nil {
		log.Println("Grupo no encontrado...")
		return models.Group{}, false
	}

	log.Println("Grupo encontrado exitosamente...")
	return group, true
}

func DeleteGroup(db *gorm.DB, id int) (string, bool) {
	var group models.Group
	db.First(&group, id)
	if group == (models.Group{}) {
		log.Println("Grupo no encontrado...")
		return "Grupo no encontrado...", false
	}

	groupDelete := db.Delete(&group)
	if groupDelete.Error != nil {
		log.Println("Error al eliminar el grupo")
		return "Error al eliminar el grupo", false
	}
	log.Println("Grupo eliminado exitosamente...")
	return "Grupo eliminado exitosamente...", true
}

func UpdateGroup(db *gorm.DB, group models.Group) (string, bool) {
	groupUpdate := db.Save(&group)
	if groupUpdate.Error != nil {
		log.Println("Error al actualizar el grupo")
		return "Error al actualizar el grupo", false
	}
	log.Println("Grupo actualizado exitosamente...")
	return "Grupo actualizado exitosamente...", true
}
