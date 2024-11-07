package models

type Section struct {
	ID           uint `gorm:"primaryKey"`
	ShowID       uint `gorm:"column:show_id;not null" json:"showId"`
	Name         string
	Price        float64 `gorm:"not null"`
	Availability uint
	ColumnStart  string `gorm:"null;default:null"`
	ColumnEnd    string `gorm:"null;default:null"`
	RowStart     uint   `gorm:"null;default:null"`
	RowEnd       uint   `gorm:"null;default:null"`
	Show         Show
}
