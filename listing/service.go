package listing

import (
	"time"
)

type Service interface {
	GetClasses() []Class
	GetClass(date time.Time) (Class, error)
	GetBookings() []Booking
}

type Repository interface {
	GetAllClasses() []Class
	GetClassForDate(date time.Time) (Class, error)
	GetAllBookings() []Booking
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetClasses() []Class {
	return s.repository.GetAllClasses()
}

func (s *service) GetClass(date time.Time) (Class, error) {
	return s.repository.GetClassForDate(date)
}

func (s *service) GetBookings() []Booking {
	return s.repository.GetAllBookings()
}
