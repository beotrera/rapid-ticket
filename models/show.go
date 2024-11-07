package models

import (
	"time"

	"gorm.io/gorm"
)

type Show struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	PlaceID     uint      `gorm:"column:place_id;not null" json:"placeId"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `gorm:"type:varchar(100);not null" json:"description"`
	Date        time.Time `gorm:"type:date;not null"`
	Place       Place
	Sections    []Section `gorm:"foreignKey:ShowID"`
}

