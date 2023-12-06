package models

type Task struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Title        string `gorm:"not null" json:"title"`
	Description  string `gorm:"not null" json:"description"`
	State        string `gorm:"not null" json:"state"`
	Group        Group  `gorm:"foreignKey:ID" json:"group_id"`
	TaskAssigned []Task `gorm:"many2many:task_assigned" json:"task_assigned"`
}
