package main

import (
	"log"
	"net/http"

	"school_management/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRoutes
	routes.RegisterSchoolRoutes(r)
	http.Handle("/", r)
	println("server is starting")
	log.Fatal(http.ListenAndServe("localhost:8009", r))

}
