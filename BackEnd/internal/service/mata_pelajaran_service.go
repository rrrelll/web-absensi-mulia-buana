package service

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
)

type MapelService struct {
	repo repository.MapelRepo
}

func NewMapelService(r repository.MapelRepo) *MapelService {
	return &MapelService{r}
}

func (s *MapelService) Create(name string) error {
	return s.repo.Create(&model.MataPelajaran{Name: name})
}

func (s *MapelService) GetAll() ([]model.MataPelajaran, error) {
	return s.repo.FindAll()
}
