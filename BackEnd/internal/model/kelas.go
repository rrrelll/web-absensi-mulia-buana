package model

type Kelas struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	JurusanID uint
}

func (Kelas) TableName() string {
	return "kelas"
}
