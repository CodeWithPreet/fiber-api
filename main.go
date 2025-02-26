package main

import (
	"log"

	"github.com/CodeWithPreet/fiber-api/database"
	"github.com/gofiber/fiber/v2"
)


func welcome(c *fiber.Ctx) error  {
	return c.SendString("welcome to my API  ....")
}
func main() {

	database.ConnectDB()
	app := fiber.New()
	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}