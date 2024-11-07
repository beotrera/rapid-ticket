package seeders

import (
	"encoding/json"
	"fmt"
	"os"

	"meli/models"

	"gorm.io/gorm"
)

func SectionSeed(db *gorm.DB) error {

	fileData, err := os.ReadFile("seeders/data/section.json")
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	var sections []models.Section
	err = json.Unmarshal(fileData, &sections)
	if err != nil {
		return fmt.Errorf("error parsing JSON file: %w", err)
	}

	for _, section := range sections {
		result := db.FirstOrCreate(&section)
		if result.Error != nil {
			return fmt.Errorf("error inserting section: %w", result.Error)
		}
	}

	return nil
}
