package main

import (
	"fmt"
	"log"
	"school_management/config"
	"school_management/v1/account"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	_, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	}
	account.AccountRoutes(app)
	log.Fatal(app.Listen(":8009"))
	println("server is starting")
}
