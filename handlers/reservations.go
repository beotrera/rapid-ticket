package handlers

import (
	"errors"
	"fmt"
	"meli/models"
	"meli/response"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReservationRequest struct {
	ShowID uint             `json:"showId"`
	DNI    string           `json:"dni"`
	Name   string           `json:"name"`
	Seats  []SectionRequest `json:"seats,omitempty"`
}

type SectionRequest struct {
	SectionID uint   `json:"sectionId"`
	Seat      string `json:"seat,omitempty"`
}

func CreateReservation(c *fiber.Ctx) error {
	var req ReservationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	
	reservation, err := Reserve(db, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(reservation)
}

func Reserve(db *gorm.DB, req ReservationRequest) (response.Ticket, error) {
	var totalPrice float64
	var seats []string
	var show models.Show


	err := db.Transaction(func(transaction *gorm.DB) error {
		if err := transaction.Where("id = ?", req.ShowID).First(&show).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("show %d does not exist", req.ShowID)
			}
			return fmt.Errorf("error getting show: %v", err)
		}

		for _, section := range req.Seats {
			var sectionData models.Section

			if err := transaction.Where("show_id = ? AND id = ?", req.ShowID, section.SectionID).First(&sectionData).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("section%d not found", section.SectionID)
				}

				return fmt.Errorf("error getting section: %v", err)
			}

			if sectionData.Availability == 0 {
				return fmt.Errorf("error creating reservation: insufficient availability")
			}

			totalPrice += sectionData.Price

			if sectionData.ColumnStart != "" && sectionData.ColumnEnd != "" {
				availableSeats := GetSeats(sectionData.ColumnStart, sectionData.ColumnEnd, sectionData.RowStart, sectionData.RowEnd, req.ShowID)

				if len(availableSeats) == 0 {
					return fmt.Errorf("error creating reservation: insufficient seats")
				}

				if !contains(availableSeats, section.Seat) {
					return fmt.Errorf("error creating reservation: the seat is already reserved")
				}

				seats = append(seats, section.Seat)
			}

			if err := transaction.Model(&models.Section{}).
				Where("id = ?", sectionData.ID).
				UpdateColumn("availability", gorm.Expr("availability - ?", 1)).Error; err != nil {
				return fmt.Errorf("error updating availability: %v", err)
			}
		}

		seatsStr := strings.Join(seats, ",")

		newReservation := models.Reservation{
			DNI:        req.DNI,
			Name:       req.Name,
			ShowID:     req.ShowID,
			Seats:      seatsStr,
			TotalPrice: totalPrice,
		}

		if err := transaction.Create(&newReservation).Error; err != nil {
			return fmt.Errorf("error creating reservation: %v", err)
		}

		return nil
	})


	ticket := response.Ticket{
		Name:       req.Name,
		DNI:        req.DNI,
		ShowName:   show.Title,
		Date:       show.Date,
		TotalPrice: totalPrice,
		Seats:      seats,
	}

	if err != nil {
		return response.Ticket{}, err
	}

	return ticket, err
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
