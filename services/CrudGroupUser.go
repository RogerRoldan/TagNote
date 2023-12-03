package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateGroupUser(groupUser *models.GroupUser, db *gorm.DB) {
	db.Create(&groupUser)
	log.Println("Usuario agregado al grupo exitosamente...")
}

func ReadGroupUser(groupUser *models.GroupUser, db *gorm.DB) {
	db.Find(&groupUser)
	log.Println("Usuario encontrado exitosamente...")
}

func DeleteGroupUser(groupUser *models.GroupUser, db *gorm.DB) {
	db.Delete(&groupUser)
	log.Println("Usuario eliminado del grupo exitosamente...")
}

func UpdateGroupUser(groupUser *models.GroupUser, db *gorm.DB) {
	db.Save(&groupUser)
	log.Println("Usuario actualizado exitosamente...")
}
