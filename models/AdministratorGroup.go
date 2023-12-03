package models

type AdministratorGroup struct {
	ID      uint `gorm:"primaryKey"`
	IDGroup int  `gorm:"not null"`
	IDUser  int  `gorm:"not null"`
}
