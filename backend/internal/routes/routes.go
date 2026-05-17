package routes

import (
	"fleetify/internal/config"
	"fleetify/internal/handlers"
	"fleetify/internal/middleware"
	"fleetify/internal/repositories"
	"fleetify/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	reportRepo := repositories.NewReportRepository(config.DB)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	api := app.Group("/api")

	api.Get("/reports", reportHandler.GetAll)

	api.Post(
		"/reports",
		middleware.RBACMiddleware("SA"),
		reportHandler.CreateReport,
	)

	api.Patch(
		"/reports/:id/approve",
		middleware.RBACMiddleware("APPROVAL"),
		reportHandler.ApproveReport,
	)

	api.Patch(
		"/reports/:id/complete",
		middleware.RBACMiddleware("SA"),
		reportHandler.CompleteReport,
	)
}