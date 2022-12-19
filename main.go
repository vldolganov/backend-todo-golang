package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"serverForMyTodo/apis"
	"serverForMyTodo/database"
)

func main() {
	app := fiber.New()

	database.InitConnection()

	apis.InitRouter(app)
	log.Fatal(app.Listen(":5000"))

}
