package handler

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/service"
	"WebAbsensiMuliaBuana/BackEnd/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	type req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON("invalid request")
	}

	err := h.service.Register(body.Name, body.Email, body.Password, body.Role)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("register success")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type req struct {
		Email    string
		Password string
	}

	var body req
	c.BodyParser(&body)

	user, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(401).JSON("login failed")
	}

	token, _ := jwt.GenerateToken(user.ID, user.Role)

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}
