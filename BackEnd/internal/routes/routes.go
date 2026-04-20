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
	guruMapelKelasHandler *handler.GuruMapelKelasHandler,
	mapelHandler *handler.MapelHandler,
	absensiGuruHandler *handler.AbsensiGuruHandler,
) {

	// ========================
	// PUBLIC ROUTES
	// ========================
	app.Post("/register", auth.Register)
	app.Post("/login", auth.Login)

	// ========================
	// PROTECTED
	// ========================
	api := app.Group("/api", middleware.AuthMiddleware())

	// ========================
	// ROLE: GURU
	// ========================
	guru := api.Group("/guru", middleware.RoleMiddleware("guru"))

	guru.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello guru")
	})

	guru.Get("/mapel-kelas", guruMapelKelasHandler.GetMy)

	// ========================
	// ROLE: ADMIN
	// ========================
	admin := api.Group("/admin", middleware.RoleMiddleware("admin"))

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello admin")
	})

	// CRUD ADMIN
	admin.Post("/jurusan", jurusanHandler.Create)
	admin.Get("/jurusan", jurusanHandler.GetAll)

	admin.Post("/kelas", kelasHandler.Create)
	admin.Get("/kelas", kelasHandler.GetAll)

	admin.Post("/mapel", mapelHandler.Create)
	admin.Get("/mapel", mapelHandler.GetAll)

	admin.Post("/assign-siswa", siswaKelasHandler.Assign)

	admin.Post("/assign-guru", guruMapelKelasHandler.Assign) // ✅ tambah

	guru.Get("/mapel-kelas", guruMapelKelasHandler.GetMy)

	guru.Post("/absen", absensiGuruHandler.Absen)
}
