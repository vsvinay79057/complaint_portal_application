package repository

import (
	"complaint_portal/models"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComplaintRepo interface {
	Create(c *models.ComplaintModel) error
	GetByID(id string) (*models.ComplaintModel, error)
	GetAllByUser(userID string) ([]models.ComplaintModel, error)
	GetAll() ([]models.ComplaintModel, error)
	MarkResolved(id string) error
}

type complaintRepo struct {
	db *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) ComplaintRepo {
	return &complaintRepo{db: db}
}

func (r *complaintRepo) Create(c *models.ComplaintModel) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return r.db.Create(c).Error
}

func (r *complaintRepo) GetByID(id string) (*models.ComplaintModel, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var cp models.ComplaintModel
	if err := r.db.Preload("User").First(&cp, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cp, nil
}

func (r *complaintRepo) GetAllByUser(userID string) ([]models.ComplaintModel, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	var list []models.ComplaintModel
	if err := r.db.Where("user_id = ?", uid).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *complaintRepo) GetAll() ([]models.ComplaintModel, error) {
	var list []models.ComplaintModel
	if err := r.db.Preload("User").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *complaintRepo) MarkResolved(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	var c models.ComplaintModel
	if err := r.db.First(&c, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	now := time.Now()
	c.Resolved = true
	c.ResolvedAt = &now
	return r.db.Save(&c).Error
}
