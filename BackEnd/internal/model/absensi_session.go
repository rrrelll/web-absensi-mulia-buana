package model

import "time"

type AbsensiSession struct {
	ID          uint `gorm:"primaryKey"`
	GuruID      uint
	KelasID     uint
	MapelID     uint
	QRToken     string
	ExpiredAt   time.Time
	Latitude    float64
	Longitude   float64
	RadiusMeter int
	CreatedAt   time.Time
}

func (AbsensiSession) TableName() string {
	return "absensi_session"
}
