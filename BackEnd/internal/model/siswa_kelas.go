package model

type SiswaKelas struct {
	ID      uint `gorm:"primaryKey"`
	SiswaID uint
	KelasID uint
}

func (SiswaKelas) TableName() string {
	return "siswa_kelas"
}
