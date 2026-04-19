package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type KelasRepository interface {
	Create(k *model.Kelas) error
	FindAll() ([]model.Kelas, error)
}

type kelasRepo struct {
	db *gorm.DB
}

func NewKelasRepo(db *gorm.DB) KelasRepository {
	return &kelasRepo{db}
}

func (r *kelasRepo) Create(k *model.Kelas) error {
	return r.db.Create(k).Error
}

func (r *kelasRepo) FindAll() ([]model.Kelas, error) {
	var data []model.Kelas
	err := r.db.Find(&data).Error
	return data, err
}
