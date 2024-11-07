package response

import "time"

type ShowResponse struct {
	ShowID      uint          `json:"showId"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Date        time.Time     `json:"date"`
	Sections    []SectionInfo `json:"sections"`
}

type SectionInfo struct {
	SectionID      uint     `json:"sectionId"`
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	ColumnStart    string   `json:"column_start,omitempty"`
	ColumnEnd      string   `json:"column_end,omitempty"`
	RowStart       uint     `json:"row_start,omitempty"`
	RowEnd         uint     `json:"row_end,omitempty"`
	AvailableSeats []string `json:"availableSeats"`
}
