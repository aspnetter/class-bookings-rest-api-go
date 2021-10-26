package adding

import (
	"errors"

	"github.com/aspnetter/go-glofox-api/listing"
)

var ErrDuplicate = errors.New("Class for this date already exists")

type Service interface {
	AddClass(...Class) error
}

type Repository interface {
	AddClass(Class) error
	GetAllClasses() []listing.Class
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddClass(c ...Class) error {

	existingClasses := s.repository.GetAllClasses()
	for _, cc := range c {
		for _, e := range existingClasses {
			if cc.Name == e.Name {
				return ErrDuplicate
			}
		}
	}

	for _, class := range c {
		_ = s.repository.AddClass(class)
	}

	return nil
}
