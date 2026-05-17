package handlers

import (
	"fleetify/internal/dto"
	"fleetify/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	Service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{Service: service}
}

func (h *ReportHandler) CreateReport(c *fiber.Ctx) error {

	var req dto.CreateReportRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userID := c.Locals("user_id").(uint)

	err := h.Service.CreateReport(req, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "report created",
	})
}

func (h *ReportHandler) ApproveReport(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := h.Service.ApproveReport(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "report approved",
	})
}

func (h *ReportHandler) CompleteReport(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	proofPhoto := c.FormValue("proof_photo")

	err := h.Service.CompleteReport(uint(id), proofPhoto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "report completed",
	})
}

func (h *ReportHandler) GetAll(c *fiber.Ctx) error {

	reports, err := h.Service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(reports)
}