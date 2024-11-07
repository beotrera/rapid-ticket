package response

import "time"

type Ticket struct {
	Name       string    `json:"name"`
	DNI        string    `json:"dni"`
	ShowName   string    `json:"showName"`
	Date       time.Time `json:"date"`
	Seats      []string  `gorm:"type:TEXT"`
	TotalPrice float64
}
