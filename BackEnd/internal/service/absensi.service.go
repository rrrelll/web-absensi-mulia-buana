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

	// 🔥 VALIDASI EXPIRED
	if time.Now().After(session.ExpiredAt) {
		return errors.New("QR expired")
	}

	// 🔥 VALIDASI SISWA ADA DI KELAS
	isValid, err := s.repo.IsSiswaInKelas(session.KelasID, siswaID)
	if err != nil {
		return err
	}

	if !isValid {
		return errors.New("anda bukan bagian dari kelas ini")
	}

	// 🔥 CEK SUDAH ABSEN
	exists, err := s.repo.CheckAlreadyAbsen(session.ID, siswaID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("sudah absen")
	}

	// 🔥 VALIDASI GEO
	d := utils.CalculateDistance(lat, lon, session.Latitude, session.Longitude)

	if d > float64(session.RadiusMeter) {
		return errors.New("anda di luar area")
	}

	// ✅ SIMPAN ABSENSI
	return s.repo.CreateAbsensi(&model.AbsensiSiswa{
		SessionID:  session.ID,
		SiswaID:    siswaID,
		WaktuAbsen: time.Now(),
		Status:     "hadir",
	})
}

// =========================
// AUTO GENERATE ALPA
// =========================
func (s *AbsensiService) GenerateAlpa(sessionID uint) error {

	session, err := s.repo.GetSessionByID(sessionID)
	if err != nil {
		return err
	}

	// 🔥 VALIDASI
	if time.Now().Before(session.ExpiredAt) {
		return errors.New("absensi masih berlangsung")
	}

	students, err := s.repo.GetSiswaByKelas(session.KelasID)
	if err != nil {
		return err
	}

	absens, err := s.repo.GetAbsensiBySession(sessionID)
	if err != nil {
		return err
	}

	absenMap := make(map[uint]bool)
	for _, a := range absens {
		absenMap[a.SiswaID] = true
	}

	for _, student := range students {

		if !absenMap[student.SiswaID] {

			exists, err := s.repo.CheckAlreadyAbsen(sessionID, student.SiswaID)
			if err != nil {
				return err
			}

			if !exists {
				err := s.repo.CreateAbsensi(&model.AbsensiSiswa{
					SessionID:  sessionID,
					SiswaID:    student.SiswaID,
					WaktuAbsen: session.ExpiredAt,
					Status:     "alpa",
				})

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// =========================
// LAPORAN
// =========================
func (s *AbsensiService) GetLaporan(sessionID uint) ([]model.AbsensiSiswa, error) {
	return s.repo.GetAbsensiBySession(sessionID)
}

// =========================
// UPDATE STATUS (GURU ONLY)
// =========================
func (s *AbsensiService) UpdateStatus(userRole string, absensiID uint, status string) error {

	// 🔥 proteksi role
	if userRole != "guru" {
		return errors.New("akses ditolak")
	}

	// validasi status
	validStatus := map[string]bool{
		"hadir": true,
		"alpa":  true,
		"izin":  true,
		"sakit": true,
	}

	if !validStatus[status] {
		return errors.New("status tidak valid")
	}

	// 🔥 optional rule (biar realistis)
	if status == "hadir" {
		return errors.New("tidak bisa ubah ke hadir")
	}

	return s.repo.UpdateStatus(absensiID, status)
}

func (s *AbsensiService) GetSummary(sessionID uint) (map[string]int64, error) {
	return s.repo.GetSummary(sessionID)
}
