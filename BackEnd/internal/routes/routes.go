package routes

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, auth *handler.AuthHandler) {
	app.Post("/register", auth.Register)
	app.Post("/login", auth.Login)
}