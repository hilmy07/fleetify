package main

import (
	"fleetify/internal/config"
	"fleetify/internal/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.MasterItem{}, &models.ReportItem{}, &models.MaintenanceReport{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.SendString("Backend is running well!")
	})

	app.Listen(":" + os.Getenv("PORT"))
}