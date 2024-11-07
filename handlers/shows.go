package handlers

import (
	"fmt"
	"meli/models"
	"meli/response"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ShowFilter struct {
	PriceMin  *float64
	PriceMax  *float64
	DateFrom  *time.Time
	DateTo    *time.Time
	OrderBy   *string
	OrderType *string
}

func ListShows(c *fiber.Ctx) error {
	var filter ShowFilter

	if priceMinStr := c.Query("priceMin"); priceMinStr != "" {
		priceMin, err := strconv.ParseFloat(priceMinStr, 64)
		if err == nil {
			filter.PriceMin = &priceMin
		}
	}

	if priceMaxStr := c.Query("priceMax"); priceMaxStr != "" {
		priceMax, err := strconv.ParseFloat(priceMaxStr, 64)
		if err == nil {
			filter.PriceMax = &priceMax
		}
	}

	if dateFromStr := c.Query("dateFrom"); dateFromStr != "" {
		dateFrom, err := time.Parse("2006-01-02", dateFromStr)
		if err == nil {
			filter.DateFrom = &dateFrom
		}
	}

	if dateToStr := c.Query("dateTo"); dateToStr != "" {
		dateTo, err := time.Parse("2006-01-02", dateToStr)
		if err == nil {
			filter.DateTo = &dateTo
		}
	}

	if orderByStr := c.Query("orderBy"); orderByStr != "" {
		filter.OrderBy = &orderByStr
	}

	if orderTypeStr := c.Query("orderType"); orderTypeStr != "" {
		filter.OrderType = &orderTypeStr
	}

	shows, err := GetShowsWithSections(db, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting shows",
		})
	}

	return c.JSON(shows)
}

func GetShowsWithSections(db *gorm.DB, filter ShowFilter) ([]response.ShowResponse, error) {
	var shows []models.Show

	query := db.Joins("JOIN section ON section.show_id = show.id").
		Preload("Sections")

	if filter.PriceMin != nil {
		query = query.Where("section.price >= ?", *filter.PriceMin)
	}
	if filter.PriceMax != nil {
		query = query.Where("section.price <= ?", *filter.PriceMax)
	}

	if filter.DateFrom != nil {
		query = query.Where("show.date >= ?", *filter.DateFrom)
	}
	if filter.DateTo != nil {
		query = query.Where("show.date <= ?", *filter.DateTo)
	}

	if filter.OrderBy != nil {
		orderType := "asc"
		if filter.OrderType != nil && (*filter.OrderType == "asc" || *filter.OrderType == "desc") {
			orderType = *filter.OrderType
		}
		query = query.Order(fmt.Sprintf("%s %s", *filter.OrderBy, orderType))
	}

	err := query.Find(&shows).Error
	if err != nil {
		return nil, err
	}

	var result []response.ShowResponse

	showMap := make(map[int]bool)

	for _, show := range shows {
		showResponse := response.ShowResponse{
			ShowID:      show.ID,
			Title:       show.Title,
			Description: show.Description,
			Date:        show.Date,
		}

		if showMap[int(show.ID)] {
			continue
		}

		for _, section := range show.Sections {
			var availableSeats []string 

			if filter.PriceMax == nil && filter.PriceMin == nil {
				if section.ColumnStart != "" && section.ColumnEnd != "" {
					availableSeats = GetSeats(section.ColumnStart, section.ColumnEnd, section.RowStart, section.RowEnd,show.ID )
				}

				showResponse.Sections = append(showResponse.Sections, response.SectionInfo{
					SectionID:   section.ID,
					Name:        section.Name,
					Price:       section.Price,
					ColumnStart: section.ColumnStart,
					ColumnEnd:   section.ColumnEnd,
					RowStart:    section.RowStart,
					RowEnd:      section.RowEnd,
					AvailableSeats: availableSeats,
				})
			} else {
				if (filter.PriceMin == nil || section.Price >= *filter.PriceMin) &&
					(filter.PriceMax == nil || section.Price <= *filter.PriceMax) {
						if section.ColumnStart != "" && section.ColumnEnd != "" {
							availableSeats = GetSeats(section.ColumnStart, section.ColumnEnd, section.RowStart, section.RowEnd,show.ID )
						}

						
					showResponse.Sections = append(showResponse.Sections, response.SectionInfo{
						SectionID:   section.ID,
						Name:        section.Name,
						Price:       section.Price,
						ColumnStart: section.ColumnStart,
						ColumnEnd:   section.ColumnEnd,
						RowStart:    section.RowStart,
						RowEnd:      section.RowEnd,
						AvailableSeats: availableSeats,
					})

				}

			}

		}

		result = append(result, showResponse)
		showMap[int(show.ID)] = true
	}

	return result, nil
}

func GetSeats(columnStart string,
	columnEnd string,
	rowStart uint,
	rowEnd uint,
	showId uint) []string {
	var seats []string
	var reservations []models.Reservation

	db.Where("show_id = ?",showId).Find(&reservations)

	reservedSeats := make(map[string]struct{})

	for _, reservation := range reservations {
		reservedSeatsList := strings.Split(reservation.Seats, ",")
		for _, reservedSeat := range reservedSeatsList {
			reservedSeats[reservedSeat] = struct{}{}
		}
	}

	for column := columnStart; column <= columnEnd; column = string([]rune(column)[0] + 1) {

		for row := rowStart; row <= rowEnd; row++ {
			rowStr := strconv.Itoa(int(row))
			seat := fmt.Sprintf("%s%s", column, rowStr)

			if _, exists := reservedSeats[seat]; !exists {
				seats = append(seats, seat)
			}
			
		}
	}

	return seats
}
