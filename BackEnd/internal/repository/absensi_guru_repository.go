package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type AbsensiGuruRepo interface {
	Create(data *model.AbsensiGuru) error
}

type absensiGuruRepo struct {
	db *gorm.DB
}

func NewAbsensiGuruRepo(db *gorm.DB) AbsensiGuruRepo {
	return &absensiGuruRepo{db}
}

func (r *absensiGuruRepo) Create(data *model.AbsensiGuru) error {
	return r.db.Create(data).Error
}
