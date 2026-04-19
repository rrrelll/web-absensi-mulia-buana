package handler

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/service"

	"github.com/gofiber/fiber/v2"
)

type JurusanHandler struct {
	service *service.JurusanService
}

func NewJurusanHandler(s *service.JurusanService) *JurusanHandler {
	return &JurusanHandler{s}
}

func (h *JurusanHandler) Create(c *fiber.Ctx) error {
	var body struct {
		Name string
	}

	c.BodyParser(&body)

	err := h.service.Create(body.Name)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("jurusan created")
}

func (h *JurusanHandler) GetAll(c *fiber.Ctx) error {
	data, _ := h.service.GetAll()
	return c.JSON(data)
}
