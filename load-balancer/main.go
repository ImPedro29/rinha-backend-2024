package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	app.Use(proxy.Balancer(proxy.Config{
		Servers: []string{
			os.Getenv("HOST_1"),
			os.Getenv("HOST_2"),
		},
	}))

	port := "9999"
	if newPort, ok := os.LookupEnv("PORT"); ok {
		port = newPort
	}

	if err := app.Listen(":" + port); err != nil {
		log.Panicln(err)
	}
}
