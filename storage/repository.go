package storage

import (
	"math/rand"
	"time"

	"cloud.google.com/go/civil"
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

	startDate, err := civil.ParseDate(c.StartDate)

	if err != nil {
		return err
	}

	endDate, err := civil.ParseDate(c.EndDate)

	if err != nil {
		return err
	}

	r.classes = append(r.classes, Class{
		ID:        GetID(),
		Name:      c.Name,
		StartDate: startDate.In(time.UTC),
		EndDate:   endDate.In(time.UTC),
		Capacity:  c.Capacity,
	})

	return nil
}

func GetID() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10000
	return (rand.Intn(max-min) + min)
}
