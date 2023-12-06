package models

type TaskAssigned struct {
	ID   uint `gorm:"primaryKey" json:"id"`
	Task Task `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"task_id" `
	User User `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id" `
}
