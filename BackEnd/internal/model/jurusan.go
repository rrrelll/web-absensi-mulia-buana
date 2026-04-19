package model

type Jurusan struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (Jurusan) TableName() string {
	return "jurusan"
}
