package controllers

import (
	"github.com/gofiber/fiber/v2"
	"serverForMyTodo/database"
	"serverForMyTodo/database/models"
)

type RequestPayload struct {
	ID      uint   `json:"id"`
	Task    string `json:"taskText"`
	IsCheck bool   `json:"isComplete"`
}

func AddTaskController(c *fiber.Ctx) error {
	var payload RequestPayload
	var db = database.Connection
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	var todo = models.Todo{
		TaskText:   payload.Task,
		IsComplete: payload.IsCheck,
	}
	db.Create(&todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func GetAllTasks(c *fiber.Ctx) error {
	var todo = []models.Todo{}
	var db = database.Connection
	db.Order("id").Find(&todo)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func DeleteAllCheckedTasks(c *fiber.Ctx) error {
	var todo = models.Todo{}
	var db = database.Connection
	db.Where("is_complete = ?", true).Delete(&todo)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func CheckAllTasks(c *fiber.Ctx) error {
	var payload RequestPayload
	var db = database.Connection
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	var todo = models.Todo{
		IsComplete: payload.IsCheck,
	}
	db.Model(&todo).Where("is_complete = ?", !todo.IsComplete).Update("is_complete", todo.IsComplete)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func DeleteTask(c *fiber.Ctx) error {
	var todo = models.Todo{}
	var db = database.Connection
	param := struct {
		ID uint `params:"id"`
	}{}
	c.ParamsParser(&param)
	db.Where("id = ?", param.ID).Delete(&todo)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func EditTaskText(c *fiber.Ctx) error {
	param := struct {
		ID uint `params:"id"`
	}{}
	var db = database.Connection
	c.ParamsParser(&param)
	var payload RequestPayload
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	var todo = models.Todo{
		ID:       param.ID,
		TaskText: payload.Task,
	}
	db.Model(&todo).Where("id = ?", param.ID).Update("task_text", todo.TaskText)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func CheckTask(c *fiber.Ctx) error {
	param := struct {
		ID uint `params:"id"`
	}{}
	var db = database.Connection
	c.ParamsParser(&param)
	var payload RequestPayload
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	var todo = models.Todo{
		IsComplete: payload.IsCheck,
	}
	db.Model(&todo).Where("id = ?", param.ID).Update("is_complete", todo.IsComplete)
	return c.Status(fiber.StatusOK).JSON(todo)
}
