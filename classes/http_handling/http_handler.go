package http_handling

import (
	"encoding/json"
	"net/http"

	"github.com/aspnetter/go-glofox-api/classes/services"
)

func GetClasses(s services.ClassService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		classesList := s.GetClasses()
		classesResponse := ParseFromEntities(classesList)

		json.NewEncoder(w).Encode(classesResponse)
	}
}

func AddClass(s services.ClassService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var newClassRequest ClassesRequestBody
		err := decoder.Decode(&newClassRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newClass, err := GetEntity(newClassRequest)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest)
		}

		s.AddClass(newClass)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New class added.")
	}
}
