package models

type TaskAssigned struct {
	ID   uint `gorm:"primaryKey"`
	Task Task `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User User `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
