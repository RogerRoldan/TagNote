package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateGroupUser(db *gorm.DB, groupUser models.GroupUser) (string, bool) {
	groupUserCreate := db.Create(&groupUser)
	if groupUserCreate.Error != nil {
		log.Println("Error al crear el grupo")
		return "Error al crear el grupo", false
	}
	log.Println("Grupo creado exitosamente...")
	return "Grupo creado exitosamente...", true
}

func GetGroupUsers(db *gorm.DB) ([]models.GroupUser, bool) {
	var groupUsers []models.GroupUser
	groupUsersList := db.Find(&models.GroupUser{}).Preload("User").Find(&groupUsers)
	if groupUsersList.Error != nil {
		log.Println("Error al obtener los grupos")
		return []models.GroupUser{}, false
	}

	log.Println("Grupos encontrados exitosamente...")
	return groupUsers, true
}

func GetGroupUserId(db *gorm.DB, id int) (models.GroupUser, bool) {
	var groupUser models.GroupUser
	response := db.First(&groupUser, id)
	if response.Error != nil {
		log.Println("Grupo no encontrado...")
		return models.GroupUser{}, false
	}

	log.Println("Grupo encontrado exitosamente...")
	return groupUser, true
}

func DeleteGroupUser(db *gorm.DB, id int) (string, bool) {
	var groupUser models.GroupUser
	var aux models.GroupUser = models.GroupUser{}
	db.First(&groupUser, id)
	if groupUser.ID == (aux.ID) {
		log.Println("Grupo no encontrado...")
		return "Grupo no encontrado...", false
	}

	groupUserDelete := db.Delete(&groupUser)
	if groupUserDelete.Error != nil {
		log.Println("Error al eliminar el grupo")
		return "Error al eliminar el grupo", false
	}
	log.Println("Grupo eliminado exitosamente...")
	return "Grupo eliminado exitosamente...", true
}

func GetGroupUsersByUserId(db *gorm.DB, id int) ([]models.GroupUser, bool) {
	var groupUsers []models.GroupUser
	groupUsersList := db.Where("user_id = ?", id).Find(&models.GroupUser{}).Preload("User").Find(&groupUsers)
	if groupUsersList.Error != nil {
		log.Println("Error al obtener los grupos")
		return []models.GroupUser{}, false
	}

	log.Println("Grupos encontrados exitosamente...")
	return groupUsers, true
}

func GetGroupUsersByGroupId(db *gorm.DB, id int) ([]models.GroupUser, bool) {
	var groupUsers []models.GroupUser
	groupUsersList := db.Where("group_id = ?", id).Find(&models.GroupUser{}).Preload("User").Find(&groupUsers)
	if groupUsersList.Error != nil {
		log.Println("Error al obtener los grupos")
		return []models.GroupUser{}, false
	}

	log.Println("Grupos encontrados exitosamente...")
	return groupUsers, true
}

func UpdateGroupUser(db *gorm.DB, groupUser models.GroupUser) (string, bool) {
	groupUserM := db.Save(&groupUser)
	if groupUserM.Error != nil {
		log.Println("Error al actualizar el grupo")
		return "Error al actualizar el grupo", false
	}
	log.Println("Grupo actualizado exitosamente...")
	return "Grupo actualizado exitosamente...", true
}
