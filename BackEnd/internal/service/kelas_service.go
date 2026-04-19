package service

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
)

type KelasService struct {
	repo repository.KelasRepository
}

func NewKelasService(r repository.KelasRepository) *KelasService {
	return &KelasService{r}
}

func (s *KelasService) Create(name string, jurusanID uint) error {
	return s.repo.Create(&model.Kelas{
		Name:      name,
		JurusanID: jurusanID,
	})
}

func (s *KelasService) GetAll() ([]model.Kelas, error) {
	return s.repo.FindAll()
}
