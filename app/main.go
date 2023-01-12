package main

import (
	"log"
	// "net/http"
	"github.com/gofiber/fiber/v2"
	"school_management/config"
	// "github.com/gorilla/mux"
)

func main() {
	app := fiber.New()
	config.Connect()
	// r := mux.NewRouter()
	// routes.RegisterSchoolRoutes(r)
	// http.Handle("/", r)

	// log.Fatal(http.ListenAndServe("localhost:8009", r))
	log.Fatal(app.Listen(":8009"))
	println("server is starting")

}
