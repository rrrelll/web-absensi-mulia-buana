package model

import "time"

type TeacherInvitation struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Token     string
	ExpiredAt time.Time
	Status    string
	CreatedAt time.Time
}

func (TeacherInvitation) TableName() string {
	return "teacher_invitations"
}
