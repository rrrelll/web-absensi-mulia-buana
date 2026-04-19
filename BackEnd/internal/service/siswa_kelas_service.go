package service

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
)

type SiswaKelasService struct {
	repo repository.SiswaKelasRepository
}

func NewSiswaKelasService(r repository.SiswaKelasRepository) *SiswaKelasService {
	return &SiswaKelasService{r}
}

func (s *SiswaKelasService) Assign(siswaID, kelasID uint) error {
	return s.repo.Assign(&model.SiswaKelas{
		SiswaID: siswaID,
		KelasID: kelasID,
	})
}
