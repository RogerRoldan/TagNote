package models

type Task struct {
	ID           uint   `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Title        string `gorm:"not null" json:"title"`
	Description  string `gorm:"not null" json:"description"`
	State        string `gorm:"not null" json:"state"`
	Group        Group  `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group_id"`
	TaskAssigned []Task `gorm:"many2many:task_assigned" json:"task_assigned"`
}
