package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type MapelRepo interface {
	Create(data *model.MataPelajaran) error
	FindAll() ([]model.MataPelajaran, error)
}

type mapelRepo struct {
	db *gorm.DB
}

func NewMapelRepo(db *gorm.DB) MapelRepo {
	return &mapelRepo{db}
}

func (r *mapelRepo) Create(data *model.MataPelajaran) error {
	return r.db.Create(data).Error
}

func (r *mapelRepo) FindAll() ([]model.MataPelajaran, error) {
	var result []model.MataPelajaran
	err := r.db.Find(&result).Error
	return result, err
}
