package model

import "time"

type AbsensiSiswa struct {
	ID         uint `gorm:"primaryKey"`
	SessionID  uint
	SiswaID    uint
	WaktuAbsen time.Time
	Status     string
}

func (AbsensiSiswa) TableName() string {
	return "absensi_siswa"
}
