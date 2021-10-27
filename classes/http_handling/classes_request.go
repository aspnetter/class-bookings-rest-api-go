package http_handling

import (
	"time"

	"cloud.google.com/go/civil"
	"github.com/aspnetter/go-glofox-api/classes/storage"
)

type ClassesRequestBody struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

func GetEntity(classesRequest ClassesRequestBody) (storage.Class, error) {
	startDate, err := civil.ParseDate(classesRequest.StartDate)

	if err != nil {
		return storage.Class{}, err
	}

	endDate, err := civil.ParseDate(classesRequest.EndDate)

	if err != nil {
		return storage.Class{}, err
	}

	return storage.Class{
		Name:      classesRequest.Name,
		StartDate: startDate.In(time.UTC),
		EndDate:   endDate.In(time.UTC),
		Capacity:  classesRequest.Capacity,
	}, nil
}
