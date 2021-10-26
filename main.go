package main

import (
	"log"
	"net/http"

	"github.com/aspnetter/go-glofox-api/adding"
	myhttp "github.com/aspnetter/go-glofox-api/http"
	"github.com/aspnetter/go-glofox-api/listing"
	"github.com/aspnetter/go-glofox-api/storage"
)

func main() {

	repo := new(storage.Repository)
	adder := adding.NewService(repo)
	lister := listing.NewService(repo)

	router := myhttp.Handler(adder, lister)

	log.Fatal(http.ListenAndServe(":8080", router))
}
