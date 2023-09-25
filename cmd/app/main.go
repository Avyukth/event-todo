package main

import (
	"log"

	"event-todo/pkg/api"
	"event-todo/pkg/application"
	"event-todo/pkg/infrastructure/eventstore"
	"event-todo/pkg/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize Event Store
	eventStore, err := eventstore.NewEventStore()
	if err != nil {
		log.Fatalf("Failed to initialize event store: %v", err)
	}

	// Initialize Read Model Repository
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize Command Handlers
	commandHandlers := application.NewCommandHandlers(eventStore)

	// Initialize Query Handlers
	queryHandlers := application.NewQueryHandlers(repo)

	// Initialize API Handlers
	apiHandlers := api.NewHandlers(commandHandlers, queryHandlers)

	// Set up Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Routes
	app.Post("/todo", apiHandlers.CreateTodoHandler)
	app.Get("/todo/:id", apiHandlers.GetTodoHandler)
	app.Get("/todos", apiHandlers.ListTodosHandler)

	// Start the Fiber app
	log.Fatal(app.Listen(":3000"))
}
