package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type JurusanRepository interface {
	Create(jurusan *model.Jurusan) error
	FindAll() ([]model.Jurusan, error)
}

type jurusanRepo struct {
	db *gorm.DB
}

func NewJurusanRepo(db *gorm.DB) JurusanRepository {
	return &jurusanRepo{db}
}

func (r *jurusanRepo) Create(j *model.Jurusan) error {
	return r.db.Create(j).Error
}

func (r *jurusanRepo) FindAll() ([]model.Jurusan, error) {
	var data []model.Jurusan
	err := r.db.Find(&data).Error
	return data, err
}
