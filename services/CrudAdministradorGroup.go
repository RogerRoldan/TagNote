package services

import (
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"log"
)

func CreateAdministrador(db *gorm.DB, administrador models.AdministratorGroup) (string, bool) {
	administradorCreate := db.Create(&administrador)
	if administradorCreate.Error != nil {
		log.Println("Error al crear el administrador")
		return "Error al crear el administrador", false
	}
	log.Println("Administrador creado exitosamente...")
	return "Administrador creado exitosamente...", true
}

func GetAdministradors(db *gorm.DB) ([]models.AdministratorGroup, bool) {
	var administradors []models.AdministratorGroup
	administradorsList := db.Find(&administradors)
	if administradorsList.Error != nil {
		log.Println("Error al obtener los administradors")
		return []models.AdministratorGroup{}, false
	}

	log.Println("Administradors encontrados exitosamente...")
	return administradors, true
}

func GetAdministradorId(db *gorm.DB, id int) (models.AdministratorGroup, bool) {
	var administrador models.AdministratorGroup
	response := db.First(&administrador, id)
	if response.Error != nil {
		log.Println("Administrador no encontrado...")
		return models.AdministratorGroup{}, false
	}

	log.Println("Administrador encontrado exitosamente...")
	return administrador, true
}

func DeleteAdministrador(db *gorm.DB, id int) (string, bool) {
	var administrador models.AdministratorGroup
	db.First(&administrador, id)
	if administrador == (models.AdministratorGroup{}) {
		log.Println("Administrador no encontrado...")
		return "Administrador no encontrado...", false
	}

	administradorDelete := db.Delete(&administrador)
	if administradorDelete.Error != nil {
		log.Println("Error al eliminar el administrador")
		return "Error al eliminar el administrador", false
	}
	log.Println("Administrador eliminado exitosamente...")
	return "Administrador eliminado exitosamente...", true
}

func UpdateAdministrador(db *gorm.DB, administrador models.AdministratorGroup) (string, bool) {
	var administradorOld models.AdministratorGroup
	db.First(&administradorOld, administrador.ID)
	if administradorOld == (models.AdministratorGroup{}) {
		log.Println("Administrador no encontrado...")
		return "Administrador no encontrado...", false
	}

	administradorUpdate := db.Model(&administradorOld).Updates(administrador)
	if administradorUpdate.Error != nil {
		log.Println("Error al actualizar el administrador")
		return "Error al actualizar el administrador", false
	}
	log.Println("Administrador actualizado exitosamente...")
	return "Administrador actualizado exitosamente...", true
}
