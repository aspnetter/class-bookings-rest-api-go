package http_handling

import "github.com/aspnetter/go-glofox-api/classes/storage"

type ClassesResponseBody struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

func ParseFromEntity(data storage.Class) ClassesResponseBody {
	dateFormat := "2006-January-02"

	return ClassesResponseBody{
		ID:        data.ID,
		Name:      data.Name,
		StartDate: data.StartDate.Format(dateFormat),
		EndDate:   data.EndDate.Format(dateFormat),
		Capacity:  data.Capacity,
	}
}

func ParseFromEntities(data []storage.Class) []ClassesResponseBody {
	var classesResponseBodies []ClassesResponseBody

	for _, class := range data {
		classesResponseBody := ParseFromEntity(class)
		classesResponseBodies = append(classesResponseBodies, classesResponseBody)
	}

	return classesResponseBodies
}
