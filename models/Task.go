package models

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	State       string `gorm:"not null"`
	GroupID     uint   `gorm:"not null"`
}
