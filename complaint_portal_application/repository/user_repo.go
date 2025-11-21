package repository

import (
	"complaint_portal/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *models.UserModel) error
	FindBySecret(secret string) (*models.UserModel, error)
	FindByID(id string) (*models.UserModel, error)
	FindAll() ([]models.UserModel, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.UserModel) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return r.db.Create(user).Error
}

func (r *userRepo) FindBySecret(secret string) (*models.UserModel, error) {
	var u models.UserModel
	if err := r.db.Preload("Complaints").Where("secret_code = ?", secret).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) FindByID(id string) (*models.UserModel, error) {
	var u models.UserModel
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	if err := r.db.Preload("Complaints").First(&u, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) FindAll() ([]models.UserModel, error) {
	var users []models.UserModel
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
