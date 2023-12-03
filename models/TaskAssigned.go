package models

type TaskAssigned struct {
	ID     uint `gorm:"primaryKey"`
	TaskID int  `gorm:"not null"`
	UserID int  `gorm:"not null"`
}
