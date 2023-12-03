package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateAdministratorGroup(db *gorm.DB, administratorGroup models.AdministratorGroup) {
	db.Create(&administratorGroup)
	log.Println("Administrador de grupo creado exitosamente...")
}

func ReadAdministratorGroup(db *gorm.DB, administratorGroup models.AdministratorGroup) {
	db.Find(&administratorGroup)
	log.Println("Administrador de grupo encontrado exitosamente...")
}

func DeleteAdministratorGroup(db *gorm.DB, administratorGroup models.AdministratorGroup) {
	db.Delete(&administratorGroup)
	log.Println("Administrador de grupo eliminado exitosamente...")
}

func UpdateAdministratorGroup(db *gorm.DB, administratorGroup models.AdministratorGroup) {
	db.Save(&administratorGroup)
	log.Println("Administrador de grupo actualizado exitosamente...")
}
