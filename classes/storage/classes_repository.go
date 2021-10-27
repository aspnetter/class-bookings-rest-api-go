package storage

import (
	"time"

	"github.com/aspnetter/go-glofox-api/internal"
)

type ClassesRepositoryInterface interface {
	GetAllClasses(date time.Time) []Class
	GetClassForDate() (Class, error)
	AddClass(c Class) error
}

type ClassesRepository struct {
	classes []Class
}

func (r *ClassesRepository) GetAllClasses() []Class {
	return r.classes
}

func (r *ClassesRepository) GetClassForDate(date time.Time) Class {

	for _, class := range r.classes {
		if internal.DateInRange(date, class.StartDate, class.EndDate) {
			return class
		}
	}
	return Class{}
}

func (r *ClassesRepository) AddClass(newClass Class) {
	newClass.ID = internal.GetID()
	newClass.CreatedDate = time.Now().UTC()

	r.classes = append(r.classes, newClass)
}
