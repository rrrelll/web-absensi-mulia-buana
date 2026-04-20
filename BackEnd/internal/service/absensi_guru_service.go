package service

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
	"errors"
	"fmt"
	"math"
	"time"
)

type AbsensiGuruService struct {
	repo repository.AbsensiGuruRepo
}

func NewAbsensiGuruService(r repository.AbsensiGuruRepo) *AbsensiGuruService {
	return &AbsensiGuruService{r}
}

// hitung jarak (meter)
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func (s *AbsensiGuruService) Absen(guruID uint, lat, lon float64, fotoPath string) error {

	// koordinat sekolah (contoh)
	const schoolLat = -6.341778
	const schoolLon = 106.558526
	const radius = 200 // meter

	d := distance(lat, lon, schoolLat, schoolLon)
	fmt.Println("INPUT LAT:", lat)
	fmt.Println("INPUT LON:", lon)
	fmt.Println("TARGET LAT:", schoolLat)
	fmt.Println("TARGET LON:", schoolLon)
	fmt.Println("DISTANCE:", d)

	if d > radius {
		return errors.New("anda di luar area sekolah")
	}

	return s.repo.Create(&model.AbsensiGuru{
		GuruID:     guruID,
		Tanggal:    time.Now(),
		WaktuAbsen: time.Now(),
		FotoPath:   fotoPath,
		Latitude:   lat,
		Longitude:  lon,
	})
}
