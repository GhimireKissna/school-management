package main

import (
	"fmt"
	"log"
	"school_management/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	_, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	}
	
	log.Fatal(app.Listen(":8009"))
	println("server is starting")
}
