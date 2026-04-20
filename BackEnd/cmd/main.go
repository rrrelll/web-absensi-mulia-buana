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
	// MATA PELAJARAN
	// =========================
	mapelRepo := repository.NewMapelRepo(db)
	mapelService := service.NewMapelService(mapelRepo)
	mapelHandler := handler.NewMapelHandler(mapelService)

	// =========================
	// SISWA KELAS
	// =========================
	siswaKelasRepo := repository.NewSiswaKelasRepo(db)
	siswaKelasService := service.NewSiswaKelasService(siswaKelasRepo)
	siswaKelasHandler := handler.NewSiswaKelasHandler(siswaKelasService)

	// =========================
	// GURU MAPEL KELAS
	// =========================
	guruMapelKelasRepo := repository.NewGuruMapelKelasRepo(db)
	guruMapelKelasService := service.NewGuruMapelKelasService(guruMapelKelasRepo)
	guruMapelKelasHandler := handler.NewGuruMapelKelasHandler(guruMapelKelasService)

	// =========================
	// ABSENSI GURU
	// =========================
	absensiGuruRepo := repository.NewAbsensiGuruRepo(db)
	absensiGuruService := service.NewAbsensiGuruService(absensiGuruRepo)
	absensiGuruHandler := handler.NewAbsensiGuruHandler(absensiGuruService)

	// =========================
	// ROUTES
	// =========================
	routes.SetupRoutes(
		app,
		authHandler,
		jurusanHandler,
		kelasHandler,
		siswaKelasHandler,
		guruMapelKelasHandler,
		mapelHandler,
		absensiGuruHandler,
	)

	app.Listen(":3000")
}
