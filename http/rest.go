package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aspnetter/go-glofox-api/adding"
	"github.com/aspnetter/go-glofox-api/listing"
)

func Handler(a adding.Service, l listing.Service) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink).Methods("GET")

	router.HandleFunc("/classes", getClasses(l)).Methods("GET")
	router.HandleFunc("/bookings", getBookings(l)).Methods("GET")

	//router.Handle("/classes", addClass(a)).Methods("POST")
	//router.Handle("/bookings", addBBooking(a)).Methods("POST")

	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getClasses(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "GET Classes!")

		w.Header().Set("Content-Type", "application/json")
		list := s.GetClasses()
		json.NewEncoder(w).Encode(list)
	}
}

func getBookings(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "GET Bookings!")

		w.Header().Set("Content-Type", "application/json")
		list := s.GetBookings()
		json.NewEncoder(w).Encode(list)
	}
}
