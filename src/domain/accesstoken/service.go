package accesstoken

import (
	"strings"

	"github.com/emadghaffari/bookstore_oauth-api/src/utils/errors"
)

// Repository interface for create a new repo for this service
type Repository interface {
	GetByID(string) (*AccessToken, *errors.ResError)
	Create(AccessToken) *errors.ResError
	Update(AccessToken) *errors.ResError
}

// Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.ResError)
	Create(AccessToken) *errors.ResError
	Update(AccessToken) *errors.ResError
}

type service struct {
	Repository Repository
}

// NewService func for create a new Service
func NewService(repo Repository) Service {
	return &service{
		Repository: repo,
	}
}

func (s service) GetByID(id string) (*AccessToken, *errors.ResError) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.HandlerBagRequest("invalid access token")
	}
	access, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return access, nil
}

func (s service) Create(ac AccessToken) *errors.ResError {
	if err := ac.Validate(); err != nil {
		return err
	}

	err := s.Repository.Create(ac)
	if err != nil {
		return err
	}

	return nil
}

func (s service) Update(ac AccessToken) *errors.ResError {
	if err := ac.Validate(); err != nil {
		return err
	}

	err := s.Repository.Update(ac)
	if err != nil {
		return err
	}
	return nil
}
