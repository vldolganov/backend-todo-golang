package apis

import (
	"github.com/gofiber/fiber/v2"
	"serverForMyTodo/apis/routes"
)

func InitRouter(app *fiber.App) {
	api := app.Group("/api")

	routes.TodoRoutes(api)
}
