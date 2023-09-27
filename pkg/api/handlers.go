package api

import (
	db "event-todo/internal"
	"event-todo/pkg/events"
	"event-todo/pkg/todo"
	"fmt"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	CommandHandler    *todo.CommandHandler
	ProjectionManager *events.ProjectionManager
	DB                *db.InMemoryDB
}

func (h *Handler) CreateTask(c *fiber.Ctx) error {
	// Parse request
	var command todo.CreateTaskCommand
	if err := c.BodyParser(&command); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	aggregateID := uuid.NewString()
	// Execute command to get the event
	event, err := command.Execute(h.CommandHandler, aggregateID) // Replace with actual aggregate ID
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot execute command"})
	}

    // ...
    if err := h.CommandHandler.EventStore.Save(aggregateID, event); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot save event"})
    }

	// Handle event to update the projection
	if err := h.ProjectionManager.HandleEvent(event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot handle event"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"aggregateID": aggregateID, "message": "Task created"})
}

func (h *Handler) CompleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	// Handle command
	command := &todo.CompleteTaskCommand{ID: id}
	event, err := command.Execute(h.CommandHandler, id) // Use Execute method here
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot complete task"})
	}

	// Save event to the event store
	if err := h.CommandHandler.EventStore.Save(id, event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot save event"})
	}

	// Handle event to update the projection
	if err := h.ProjectionManager.HandleEvent(event); err != nil {
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
	event, err := command.Execute(h.CommandHandler, id) // Use Execute method here
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete task"})
	}

	// Save event to the event store
	if err := h.CommandHandler.EventStore.Save(id, event); err != nil { // Fixed here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot save event"})
	}

	// Handle event to update the projection
	if err := h.ProjectionManager.HandleEvent(event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot handle event"})
	}

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Task %s deleted", id)})
}
