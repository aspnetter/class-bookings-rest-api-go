package services

import (
	"errors"
	"time"

	"github.com/aspnetter/go-glofox-api/classes/storage"
	"github.com/aspnetter/go-glofox-api/internal"
)

var ErrOverlap = errors.New("class for this date already exists")

type ClassService interface {
	AddClass(storage.Class) (storage.Class, error)
	GetClasses() []storage.Class
}

type Repository interface {
	AddClass(storage.Class) error
	GetAllClasses() []storage.Class
	GetClassForDate(time.Time) storage.Class
}

type service struct {
	repository Repository
}

func NewService(r Repository) ClassService {
	return &service{r}
}

func (s *service) GetClasses() []storage.Class {
	return s.repository.GetAllClasses()
}

func (s *service) AddClass(newClass storage.Class) (storage.Class, error) {

	allClasses := s.repository.GetAllClasses()

	for _, class := range allClasses {
		startDate := class.StartDate
		endDate := class.EndDate
		if internal.DateInRange(newClass.StartDate, startDate, endDate) || internal.DateInRange(newClass.EndDate, startDate, endDate) {
			return storage.Class{}, ErrOverlap
		}
	}

	s.repository.AddClass(newClass)

	return newClass, nil
}
