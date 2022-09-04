package movie

import (
	"MovieAPI/internal/core"
	"MovieAPI/pkg/pagination"
)

type Service interface {
	GetAll(pageOptions *pagination.Pages) ([]Movie, error)
	GetByParam(params map[string]string) (*Movie, error)
	Add(entity *Movie) error
	Update(entity *Movie) error
	Delete(id string) error
}

type service struct {
	repo Repository
}

func (s *service) GetAll(pageOptions *pagination.Pages) (movie []Movie, err error) {

	movie, err = s.repo.Find(pageOptions)
	if err != nil {
		return movie, err
	}
	return movie, err
}
func (s *service) GetByParam(params map[string]string) (*Movie, error) {
	return s.repo.FindOne(params)
}
func (s *service) Add(entity *Movie) error {

	filter := map[string]string{
		"Name":        entity.Name,
		"Description": entity.Description,
		"Type":        entity.Type,
	}

	hasItem, err := s.repo.FindOne(filter)
	if err != nil {
		return err
	}
	if hasItem != nil {
		return core.ErrConflict
	}
	return s.repo.Create(entity)
}
func (s *service) Update(entity *Movie) (err error) {

	hasItem, err := s.repo.GetById(entity.Id)
	if err != nil {
		return err
	}
	if hasItem == nil {
		return core.ErrNotFound
	}
	err = s.repo.Update(entity)
	return err
}

func (s *service) Delete(id string) (err error) {

	hasItem, err := s.repo.GetById(id)
	if err != nil {
		return err
	}
	if hasItem == nil {
		return core.ErrNotFound
	}
	err = s.repo.Delete(id)
	return err
}

// NewService creates a new toggle service.
func newService(repo Repository) Service {
	return &service{repo}
}
