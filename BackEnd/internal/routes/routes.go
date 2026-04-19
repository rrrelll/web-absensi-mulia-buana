package routes

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/handler"
	"WebAbsensiMuliaBuana/BackEnd/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	auth *handler.AuthHandler,
	jurusanHandler *handler.JurusanHandler,
	kelasHandler *handler.KelasHandler,
	siswaKelasHandler *handler.SiswaKelasHandler,
) {

	// ========================
	// PUBLIC ROUTES
	// ========================
	app.Post("/register", auth.Register)
	app.Post("/login", auth.Login)

	// ========================
	// PROTECTED (LOGIN WAJIB)
	// ========================
	api := app.Group("/api", middleware.AuthMiddleware())

	// ========================
	// ROLE: GURU
	// ========================
	guru := api.Group("/guru", middleware.RoleMiddleware("guru"))

	guru.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello guru")
	})

	// ========================
	// ROLE: ADMIN
	// ========================
	admin := api.Group("/admin", middleware.RoleMiddleware("admin"))

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello admin")
	})

	// 🔥 CRUD ADMIN
	admin.Post("/jurusan", jurusanHandler.Create)
	admin.Get("/jurusan", jurusanHandler.GetAll)

	admin.Post("/kelas", kelasHandler.Create)
	admin.Get("/kelas", kelasHandler.GetAll)

	admin.Post("/assign-siswa", siswaKelasHandler.Assign)
}
