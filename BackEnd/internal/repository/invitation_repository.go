package repository

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"

	"gorm.io/gorm"
)

type InvitationRepo interface {
	Create(data *model.TeacherInvitation) error
	FindByToken(token string) (*model.TeacherInvitation, error)
	Update(data *model.TeacherInvitation) error
}

type invitationRepo struct {
	db *gorm.DB
}

func NewInvitationRepo(db *gorm.DB) InvitationRepo {
	return &invitationRepo{db}
}

func (r *invitationRepo) Create(data *model.TeacherInvitation) error {
	return r.db.Create(data).Error
}

func (r *invitationRepo) FindByToken(token string) (*model.TeacherInvitation, error) {
	var inv model.TeacherInvitation
	err := r.db.Where("token = ?", token).First(&inv).Error
	return &inv, err
}

func (r *invitationRepo) Update(data *model.TeacherInvitation) error {
	return r.db.Save(data).Error
}
