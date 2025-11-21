package usecase

import (
	"complaint_portal/models"
	"complaint_portal/repository"
	"errors"
)

type ComplaintUsecase interface {
	Submit(userSecret string, req models.ComplaintRequest) (*models.ComplaintModel, error)
	GetAllForUser(userSecret string) ([]models.ComplaintModel, error)
	GetAllForAdmin() ([]models.ComplaintModel, error)
	ViewComplaint(userSecret, id string) (*models.ComplaintModel, error)
	ResolveComplaint(id string) error
}

type complaintUsecase struct {
	compRepo repository.ComplaintRepo
	userRepo repository.UserRepo
}

func NewComplaintUsecase(c repository.ComplaintRepo, u repository.UserRepo) ComplaintUsecase {
	return &complaintUsecase{compRepo: c, userRepo: u}
}

func (s *complaintUsecase) Submit(userSecret string, req models.ComplaintRequest) (*models.ComplaintModel, error) {
	user, err := s.userRepo.FindBySecret(userSecret)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	c := &models.ComplaintModel{
		Title:   req.Title,
		Summary: req.Summary,
		Rating:  req.Rating,
		UserID:  user.ID,
	}
	if err := s.compRepo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *complaintUsecase) GetAllForUser(userSecret string) ([]models.ComplaintModel, error) {
	user, err := s.userRepo.FindBySecret(userSecret)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return s.compRepo.GetAllByUser(user.ID.String())
}

func (s *complaintUsecase) GetAllForAdmin() ([]models.ComplaintModel, error) {
	return s.compRepo.GetAll()
}

func (s *complaintUsecase) ViewComplaint(userSecret, id string) (*models.ComplaintModel, error) {
	user, err := s.userRepo.FindBySecret(userSecret)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	c, err := s.compRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("complaint not found")
	}

	if user.IsAdmin || c.UserID == user.ID {
		return c, nil
	}
	return nil, errors.New("unauthorized to view complaint")
}

func (s *complaintUsecase) ResolveComplaint(id string) error {
	c, err := s.compRepo.GetByID(id)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("complaint not found")
	}
	return s.compRepo.MarkResolved(id)
}
