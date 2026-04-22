package service

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
	"WebAbsensiMuliaBuana/BackEnd/internal/utils"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AbsensiService struct {
	repo repository.AbsensiRepo
}

func NewAbsensiService(r repository.AbsensiRepo) *AbsensiService {
	return &AbsensiService{r}
}

// =========================
// GURU BUAT QR
// =========================
func (s *AbsensiService) CreateSession(guruID, kelasID, mapelID uint, lat, lon float64) (*model.AbsensiSession, error) {

	token := uuid.NewString()

	session := &model.AbsensiSession{
		GuruID:      guruID,
		KelasID:     kelasID,
		MapelID:     mapelID,
		QRToken:     token,
		ExpiredAt:   time.Now().Add(15 * time.Minute),
		Latitude:    lat,
		Longitude:   lon,
		RadiusMeter: 100,
	}

	err := s.repo.CreateSession(session)
	return session, err
}

// =========================
// MURID ABSEN
// =========================
func (s *AbsensiService) AbsenSiswa(token string, siswaID uint, lat, lon float64) error {

	session, err := s.repo.FindSessionByToken(token)
	if err != nil {
		return errors.New("QR tidak valid")
	}

	if time.Now().After(session.ExpiredAt) {
		return errors.New("QR expired")
	}

	if s.repo.CheckAlreadyAbsen(session.ID, siswaID) {
		return errors.New("sudah absen")
	}

	d := utils.CalculateDistance(lat, lon, session.Latitude, session.Longitude)

	if d > float64(session.RadiusMeter) {
		return errors.New("anda di luar area")
	}

	return s.repo.CreateAbsensi(&model.AbsensiSiswa{
		SessionID:  session.ID,
		SiswaID:    siswaID,
		WaktuAbsen: time.Now(),
		Status:     "hadir",
	})
}
