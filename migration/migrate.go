package migration

import (
	"github.com/ValeriiaHuza/exchange_rate_api/initializers"
	"github.com/ValeriiaHuza/exchange_rate_api/models"
	"log"
)

func AutomatedMigration() {
	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}
