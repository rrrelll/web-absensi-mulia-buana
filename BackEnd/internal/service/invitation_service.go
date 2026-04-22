package service

import (
	"WebAbsensiMuliaBuana/BackEnd/internal/model"
	"WebAbsensiMuliaBuana/BackEnd/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type InvitationService struct {
	repo     repository.InvitationRepo
	userRepo repository.UserRepository
}

func NewInvitationService(r repository.InvitationRepo, u repository.UserRepository) *InvitationService {
	return &InvitationService{r, u}
}

// ADMIN INVITE
func (s *InvitationService) Invite(email string) (*model.TeacherInvitation, error) {

	token := uuid.NewString()

	inv := &model.TeacherInvitation{
		Email:     email,
		Token:     token,
		ExpiredAt: time.Now().Add(24 * time.Hour),
		Status:    "pending",
	}

	err := s.repo.Create(inv)
	if err != nil {
		return nil, err
	}

	return inv, nil
}

// VALIDATE TOKEN
func (s *InvitationService) Validate(token string) (*model.TeacherInvitation, error) {

	inv, err := s.repo.FindByToken(token)
	if err != nil {
		return nil, errors.New("token tidak valid")
	}

	if inv.Status != "pending" {
		return nil, errors.New("token sudah digunakan")
	}

	if time.Now().After(inv.ExpiredAt) {
		return nil, errors.New("token expired")
	}

	return inv, nil
}

// REGISTER GURU
func (s *InvitationService) Register(token, name, password string) error {

	inv, err := s.Validate(token)
	if err != nil {
		return err
	}

	// 🔥 HASH PASSWORD
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     name,
		Email:    inv.Email,
		Password: string(hashedPassword),
		Role:     "guru",
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	inv.Status = "used"
	return s.repo.Update(inv)
}
