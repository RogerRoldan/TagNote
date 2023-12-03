package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateGroup(db *gorm.DB, group models.Group) {
	db.Create(&group)
	log.Println("Grupo creado exitosamente...")
}

func ReadGroup(db *gorm.DB, group models.Group) {
	db.Find(&group)
	log.Println("Grupo encontrado exitosamente...")
}

func DeleteGroup(db *gorm.DB, group models.Group) {
	db.Delete(&group)
	log.Println("Grupo eliminado exitosamente...")
}

func UpdateGroup(db *gorm.DB, group models.Group) {
	db.Save(&group)
	log.Println("Grupo actualizado exitosamente...")
}
