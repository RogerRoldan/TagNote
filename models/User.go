package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique; not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique; not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Imagen    []byte `gorm:"null"`
}
