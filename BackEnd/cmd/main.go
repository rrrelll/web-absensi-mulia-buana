package main

import (
	"WebAbsensiMuliaBuana/BackEnd/config"
	"WebAbsensiMuliaBuana/BackEnd/internal/handler"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
	"WebAbsensiMuliaBuana/BackEnd/internal/routes"
	"WebAbsensiMuliaBuana/BackEnd/internal/service"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env")
	}

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))

	app := fiber.New()

	// connect DB
	db := config.ConnectDB()

	// =========================
	// AUTH
	// =========================
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// =========================
	// JURUSAN
	// =========================
	jurusanRepo := repository.NewJurusanRepo(db)
	jurusanService := service.NewJurusanService(jurusanRepo)
	jurusanHandler := handler.NewJurusanHandler(jurusanService)

	// =========================
	// KELAS
	// =========================
	kelasRepo := repository.NewKelasRepo(db)
	kelasService := service.NewKelasService(kelasRepo)
	kelasHandler := handler.NewKelasHandler(kelasService)

	// =========================
	// SISWA KELAS
	// =========================
	siswaKelasRepo := repository.NewSiswaKelasRepo(db)
	siswaKelasService := service.NewSiswaKelasService(siswaKelasRepo)
	siswaKelasHandler := handler.NewSiswaKelasHandler(siswaKelasService)

	// =========================
	// ROUTES
	// =========================
	routes.SetupRoutes(
		app,
		authHandler,
		jurusanHandler,
		kelasHandler,
		siswaKelasHandler,
	)

	app.Listen(":3000")
}
