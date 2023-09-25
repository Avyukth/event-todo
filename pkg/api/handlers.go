package api

import (
	"event-todo/pkg/application"
	"event-todo/pkg/domain"

	"github.com/gofiber/fiber/v2"
)

// Handlers struct holds the command and query handlers
type Handlers struct {
	CommandHandlers *application.CommandHandlers
	QueryHandlers   *application.QueryHandlers
}

// NewHandlers initializes the API handlers with the given command and query handlers
func NewHandlers(ch *application.CommandHandlers, qh *application.QueryHandlers) *Handlers {
	return &Handlers{
		CommandHandlers: ch,
		QueryHandlers:   qh,
	}
}

// CreateTodoHandler handles the creation of a new Todo item
func (h *Handlers) CreateTodoHandler(c *fiber.Ctx) error {
	// Extract data from request
	var command domain.CreateTodoCommand
	if err := c.BodyParser(&command); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Handle the command
	if err := h.CommandHandlers.HandleCreateTodoCommand(&command); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Todo successfully created"})
}

// GetTodoHandler handles the retrieval of a Todo item by ID
func (h *Handlers) GetTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := h.QueryHandlers.HandleGetTodoQuery(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.JSON(todo)
}

// ListTodosHandler handles the retrieval of all Todo items
func (h *Handlers) ListTodosHandler(c *fiber.Ctx) error {
	todos, err := h.QueryHandlers.HandleListTodosQuery()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(todos)
}
