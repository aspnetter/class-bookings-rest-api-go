package http_handling

import (
	"github.com/aspnetter/go-glofox-api/classes/services"
	"github.com/gorilla/mux"
)

func AddClassesHandlers(router mux.Router) {

	classService := new(services.ClassService)
	router.HandleFunc("/classes", GetClasses(*classService)).Methods("GET")
	router.HandleFunc("/classes", AddClass(*classService)).Methods("POST")
}
