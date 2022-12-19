package routes

import (
	"github.com/gofiber/fiber/v2"
	"serverForMyTodo/apis/controllers"
)

func TodoRoutes(api fiber.Router) {
	TodoApi := api.Group("/")

	TodoApi.Post("/", controllers.AddTaskController)
	TodoApi.Get("/", controllers.GetAllTasks)
	TodoApi.Delete("/", controllers.DeleteAllCheckedTasks)
	TodoApi.Put("/", controllers.CheckAllTasks)

	TodoApi.Delete("/:id", controllers.DeleteTask)
	TodoApi.Put("/:id", controllers.EditTaskText)
	TodoApi.Put("/:id", controllers.CheckTask)

}
