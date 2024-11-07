package models

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Address  string `gorm:"type:varchar(200);not null" json:"address"`
	Capacity int    `gorm:"not null" json:"capacity"`
	Shows    []Show `gorm:"foreignKey:PlaceID"`
}
