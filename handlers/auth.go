package handlers

import (
	"encoding/base64"
	"meli/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization") 

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Credentials",
		})
	}

	if !strings.HasPrefix(authHeader, "Basic ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	encodedCredentials := authHeader[len("Basic "):]
	decoded, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "request",
		})
	}

	var user models.User
    if err := db.WithContext(c.Context()).Where("email = ?", credentials[0]).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid credentials",
        })
    }
	
	errBcrypt := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials[1]))
	if errBcrypt != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return c.Next()
}
