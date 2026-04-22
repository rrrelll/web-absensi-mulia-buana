package handler

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AbsensiHandler struct {
	service *service.AbsensiService
}

func NewAbsensiHandler(s *service.AbsensiService) *AbsensiHandler {
	return &AbsensiHandler{s}
}

// =========================
// GURU BUAT QR
// =========================
func (h *AbsensiHandler) CreateSession(c *fiber.Ctx) error {

	type req struct {
		KelasID   uint    `json:"kelas_id"`
		MapelID   uint    `json:"mapel_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	guruID := c.Locals("user_id").(uint)

	session, err := h.service.CreateSession(guruID, body.KelasID, body.MapelID, body.Latitude, body.Longitude)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(session)
}

// =========================
// MURID ABSEN
// =========================
func (h *AbsensiHandler) Absen(c *fiber.Ctx) error {

	type req struct {
		Token     string  `json:"token"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	siswaID := c.Locals("user_id").(uint)

	err := h.service.AbsenSiswa(body.Token, siswaID, body.Latitude, body.Longitude)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON("absen berhasil")
}
