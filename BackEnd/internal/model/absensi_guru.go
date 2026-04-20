package model

import "time"

type AbsensiGuru struct {
	ID         uint `gorm:"primaryKey"`
	GuruID     uint
	Tanggal    time.Time
	WaktuAbsen time.Time
	FotoPath   string
	Latitude   float64
	Longitude  float64
}

func (AbsensiGuru) TableName() string {
	return "absensi_guru"
}
