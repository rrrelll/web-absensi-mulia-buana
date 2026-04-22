package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type AbsensiRepo interface {
	CreateSession(data *model.AbsensiSession) error
	FindSessionByToken(token string) (*model.AbsensiSession, error)
	CreateAbsensi(data *model.AbsensiSiswa) error
	CheckAlreadyAbsen(sessionID, siswaID uint) bool
}

type absensiRepo struct {
	db *gorm.DB
}

func NewAbsensiRepo(db *gorm.DB) AbsensiRepo {
	return &absensiRepo{db}
}

func (r *absensiRepo) CreateSession(data *model.AbsensiSession) error {
	return r.db.Create(data).Error
}

func (r *absensiRepo) FindSessionByToken(token string) (*model.AbsensiSession, error) {
	var s model.AbsensiSession
	err := r.db.Where("qr_token = ?", token).First(&s).Error
	return &s, err
}

func (r *absensiRepo) CreateAbsensi(data *model.AbsensiSiswa) error {
	return r.db.Create(data).Error
}

func (r *absensiRepo) CheckAlreadyAbsen(sessionID, siswaID uint) bool {
	var count int64
	r.db.Model(&model.AbsensiSiswa{}).
		Where("session_id = ? AND siswa_id = ?", sessionID, siswaID).
		Count(&count)

	return count > 0
}
