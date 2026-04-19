package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type SiswaKelasRepository interface {
	Assign(s *model.SiswaKelas) error
}

type siswaKelasRepo struct {
	db *gorm.DB
}

func NewSiswaKelasRepo(db *gorm.DB) SiswaKelasRepository {
	return &siswaKelasRepo{db}
}

func (r *siswaKelasRepo) Assign(s *model.SiswaKelas) error {
	return r.db.Create(s).Error
}
