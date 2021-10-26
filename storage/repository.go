package storage

import (
	"time"

	"github.com/aspnetter/go-glofox-api/adding"
	"github.com/aspnetter/go-glofox-api/listing"
)

type Repository struct {
	classes []Class
	//bookings []Booking
}

func (r *Repository) GetAllClasses() []listing.Class {
	var result []listing.Class

	for i := range r.classes {

		class := listing.Class{
			ID:        r.classes[i].ID,
			Name:      r.classes[i].Name,
			StartDate: r.classes[i].StartDate,
			EndDate:   r.classes[i].EndDate,
			Capacity:  r.classes[i].Capacity,
		}

		result = append(result, class)
	}

	return result
}

func (r *Repository) GetClassForDate(date time.Time) (listing.Class, error) {
	var result listing.Class
	return result, nil
}

func (r *Repository) GetAllBookings() []listing.Booking {
	var result []listing.Booking
	return result
}

func (r *Repository) AddClass(c adding.Class) error {
	return nil
}
