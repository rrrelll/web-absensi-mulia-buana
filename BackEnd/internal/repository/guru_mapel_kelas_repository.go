package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type GuruMapelKelasRepo interface {
	Assign(data *model.GuruMapelKelas) error
	FindByGuru(guruID uint) ([]model.GuruMapelKelas, error)
}

type guruMapelKelasRepo struct {
	db *gorm.DB
}

func NewGuruMapelKelasRepo(db *gorm.DB) GuruMapelKelasRepo {
	return &guruMapelKelasRepo{db}
}

func (r *guruMapelKelasRepo) Assign(data *model.GuruMapelKelas) error {
	return r.db.Create(data).Error
}

func (r *guruMapelKelasRepo) FindByGuru(guruID uint) ([]model.GuruMapelKelas, error) {
	var result []model.GuruMapelKelas
	err := r.db.Where("guru_id = ?", guruID).Find(&result).Error
	return result, err
}
