package handler

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/service"

	"github.com/gofiber/fiber/v2"
)

type SiswaKelasHandler struct {
	service *service.SiswaKelasService
}

func NewSiswaKelasHandler(s *service.SiswaKelasService) *SiswaKelasHandler {
	return &SiswaKelasHandler{s}
}

func (h *SiswaKelasHandler) Assign(c *fiber.Ctx) error {
	var body struct {
		SiswaID uint
		KelasID uint
	}

	c.BodyParser(&body)

	err := h.service.Assign(body.SiswaID, body.KelasID)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("siswa assigned")
}
