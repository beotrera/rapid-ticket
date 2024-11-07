package seeders

import (
	"encoding/json"
	"fmt"
	"os"

	"meli/models"

	"gorm.io/gorm"
)

func PlaceSeed(db *gorm.DB) error {

	fileData, err := os.ReadFile("seeders/data/place.json")
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	var places []models.Place
	err = json.Unmarshal(fileData, &places)
	if err != nil {
		return fmt.Errorf("error parsing JSON file: %w", err)
	}

	for _, place := range places {
		result := db.FirstOrCreate(&place)
		if result.Error != nil {
			return fmt.Errorf("error inserting palces: %w", result.Error)
		}
	}

	return nil
}
