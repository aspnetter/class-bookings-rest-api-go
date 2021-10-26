package adding

import (
	"errors"

	"github.com/aspnetter/go-glofox-api/listing"
)

var ErrDuplicate = errors.New("Class for this date already exists")

type Service interface {
	AddClass(...Class) error
	AddSampleClasses()
	AddSampleBookings()
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

func (s *service) AddSampleClasses() {
	var defaultClasses = []Class{
		{
			Name:     "Yoga",
			Capacity: 20,
		},
		{
			Name:     "Pilates",
			Capacity: 10,
		},
	}

	for _, cc := range defaultClasses {
		_ = s.repository.AddClass(cc)
	}
}

func (s *service) AddSampleBookings() {

}
