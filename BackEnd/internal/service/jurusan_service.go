package service

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
)

type JurusanService struct {
	repo repository.JurusanRepository
}

func NewJurusanService(r repository.JurusanRepository) *JurusanService {
	return &JurusanService{r}
}

func (s *JurusanService) Create(name string) error {
	return s.repo.Create(&model.Jurusan{Name: name})
}

func (s *JurusanService) GetAll() ([]model.Jurusan, error) {
	return s.repo.FindAll()
}
