package main

import (
	"log"
	"net/http"

	"github.com/aspnetter/go-glofox-api/classes/http_handling"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	http_handling.AddClassesHandlers(*router)

	/*
		add classes and bookings handlers
	*/

	log.Fatal(http.ListenAndServe(":8080", router))
}
