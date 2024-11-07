package seeders

import (
	"encoding/json"
	"fmt"
	"os"

	"meli/models"

	"gorm.io/gorm"
)

func ShowSeed(db *gorm.DB) error {

	fileData, err := os.ReadFile("seeders/data/show.json")
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	var shows []models.Show
	err = json.Unmarshal(fileData, &shows)
	if err != nil {
		return fmt.Errorf("error parsing JSON file: %w", err)
	}

	for _, show := range shows {
		result := db.FirstOrCreate(&show)
		if result.Error != nil {
			return fmt.Errorf("error inserting show: %w", result.Error)
		}
	}

	return nil
}
