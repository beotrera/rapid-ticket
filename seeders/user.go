package seeders

import (
	"encoding/json"
	"fmt"
	"os"

	"meli/models"

	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) error {

	fileData, err := os.ReadFile("seeders/data/user.json")
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	var users []models.User
	err = json.Unmarshal(fileData, &users)
	if err != nil {
		return fmt.Errorf("error parsing JSON file: %w", err)
	}

	for _, user := range users {
		result := db.FirstOrCreate(&user)
		if result.Error != nil {
			return fmt.Errorf("error inserting user: %w", result.Error)
		}
	}

	return nil
}
