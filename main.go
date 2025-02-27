package main

import (
	"log"

	"github.com/CodeWithPreet/fiber-api/database"
	"github.com/CodeWithPreet/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)


func welcome(c *fiber.Ctx) error  {
	return c.SendString("welcome to my API  ....")
}


func setupRoutes(app *fiber.App){
	app.Get("welcome",welcome)

	// users
	routes.UserControllers(app)


}
func main() {

	database.ConnectDB()
	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}