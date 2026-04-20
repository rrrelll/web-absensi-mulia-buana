package model

type MataPelajaran struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (MataPelajaran) TableName() string {
	return "mata_pelajaran"
}
