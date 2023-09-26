package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"event-todo/pkg/events"
	"event-todo/pkg/todo"
)

type Handler struct {
	CommandHandler *todo.CommandHandler
	EventHandler   *events.EventHandler
	DB             todo.DB
}

func (h *Handler) CreateTask(c *fiber.Ctx) error {
	// Parse request
	var command todo.CreateTaskCommand
	if err := c.BodyParser(&command); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Handle command
	event, err := h.CommandHandler.HandleCreateTaskCommand(&command)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create task"})
	}

	// Handle event
	if err := h.EventHandler.HandleEvent(event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot handle event"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Task created"})
}

func (h *Handler) CompleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	// Handle command
	command := &todo.CompleteTaskCommand{ID: id}
	event, err := h.CommandHandler.HandleCompleteTaskCommand(command)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot complete task"})
	}

	// Handle event
	if err := h.EventHandler.HandleEvent(event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot handle event"})
	}

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Task %s completed", id)})
}

func (h *Handler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	// Handle command
	command := &todo.DeleteTaskCommand{ID: id}
	event, err := h.CommandHandler.HandleDeleteTaskCommand(command)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete task"})
	}

	// Handle event
	if err := h.EventHandler.HandleEvent(event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot handle event"})
	}

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Task %s deleted", id)})
}
