package usecase

import (
	"complaint_portal/models"
	"complaint_portal/repository"
	"complaint_portal/utils"
	"errors"
	"os"
)

type UserUsecase interface {
	Register(name, email string) (*models.UserModel, error)
	LoginWithSecret(secret string) (*models.UserModel, error)
	FindByID(id string) (*models.UserModel, error)
	CreateAdmin(name, email, adminKey string) (*models.UserModel, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(u repository.UserRepo) UserUsecase {
	return &userUsecase{userRepo: u}
}

func (s *userUsecase) Register(name, email string) (*models.UserModel, error) {
	secret := utils.GenerateSecretCode()
	user := &models.UserModel{
		SecretCode: secret,
		Name:       name,
		Email:      email,
		IsAdmin:    false,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userUsecase) LoginWithSecret(secret string) (*models.UserModel, error) {
	user, err := s.userRepo.FindBySecret(secret)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid secret code")
	}
	return user, nil
}

func (s *userUsecase) FindByID(id string) (*models.UserModel, error) {
	return s.userRepo.FindByID(id)
}

func (s *userUsecase) CreateAdmin(name, email, adminKey string) (*models.UserModel, error) {
	key := os.Getenv("ADMIN_CREATION_KEY")
	if key == "" || adminKey != key {
		return nil, errors.New("invalid admin creation key")
	}
	secret := utils.GenerateSecretCode()
	user := &models.UserModel{
		SecretCode: secret,
		Name:       name,
		Email:      email,
		IsAdmin:    true,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
