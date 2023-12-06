package models

type AdministratorGroup struct {
	ID    uint  `gorm:"primaryKey"`
	Group Group `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User  User  `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
