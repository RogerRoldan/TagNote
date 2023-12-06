package models

type AdministratorGroup struct {
	ID    uint  `gorm:"primaryKey" json:"id"`
	Group Group `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"id_group"`
	User  User  `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"id_user"`
}
