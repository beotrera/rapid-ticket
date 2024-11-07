package models

type UserType string

const (
	Admin     UserType = "ADMIN"
	Plataform UserType = "PLATAFORM"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	UserType UserType
}
