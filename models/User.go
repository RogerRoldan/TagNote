package models

type User struct {
	ID                 uint    `gorm:"primaryKey,autoIncrement" json:"id"`
	Username           string  `gorm:"unique; not null" json:"username"`
	Password           string  `gorm:"not null" json:"password"`
	Email              string  `gorm:"unique; not null" json:"email"`
	FirstName          string  `gorm:"not null" json:"first_name"`
	LastName           string  `gorm:"not null" json:"last_name"`
	Imagen             string  `gorm:"null" json:"imagen"`
	TaskAssigned       []Task  `gorm:"many2many:task_assigned" json:"task_assigned"`
	GroupsUser         []Group `gorm:"many2many:group_user" json:"groups_user"`
	AdministratorGroup []Group `gorm:"many2many:administrator_group" json:"administrator_group"`
}
