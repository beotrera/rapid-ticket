package models

type Reservation struct {
	ID         uint `gorm:"primaryKey"`
	DNI        string
	Name       string
	ShowID     uint   `gorm:"column:show_id;not null" json:"showId"`
	Seats      string `gorm:"type:TEXT"`
	TotalPrice float64
}
