package main

import (
	"fleetify/internal/config"
	"fleetify/internal/models"
	"fleetify/internal/routes"
	"fleetify/internal/seeder"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.MasterItem{},
		&models.MaintenanceReport{},
		&models.ReportItem{},
	)

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) {
	// 	c.SendString("Backend is running well!")
	// })

	seeder.Seed()

	routes.SetupRoutes(app)
	app.Listen(":" + os.Getenv("PORT"))
}