package model

type GuruMapelKelas struct {
	ID      uint `gorm:"primaryKey"`
	GuruID  uint
	KelasID uint
	MapelID uint
}

func (GuruMapelKelas) TableName() string {
	return "guru_mapel_kelas"
}
