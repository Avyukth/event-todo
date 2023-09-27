package main

import (
	"log"

	db "event-todo/internal"
	"event-todo/pkg/api"
	"event-todo/pkg/events"
	"event-todo/pkg/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()
	app.Use(logger.New())

	// Initialize in-memory event store
	eventStore := events.NewInMemoryEventStore()

	// Initialize in-memory database
	inMemoryDB := db.NewInMemoryDB()

	// Initialize Projection Manager
	projectionManager := todo.NewProjectionManager()

	// Initialize event handler with Projection Manager
	// eventHandler := events.NewEventHandler()
	// eventHandler.ProjectionManager = projectionManager // Assuming EventHandler has a ProjectionManager field

	// // Initialize command handlers
	commandHandler := &todo.CommandHandler{
		EventStore: eventStore,
		ProjectionManager: projectionManager, // Injecting ProjectionManager
	}

	// // Initialize API handlers and inject dependencies
	// apiHandler := &api.Handler{
	// 	CommandHandler: commandHandler,
	// 	EventHandler:   eventHandler,
	// 	DB:             inMemoryDB, // Assuming Handler has a DB field
	// }


	apiHandler	:= &api.Handler{
		CommandHandler: commandHandler,
		
		DB:             inMemoryDB, // Assuming Handler has a DB field

		ProjectionManager: projectionManager, // Injecting ProjectionManager
		// Inject other dependencies if needed
	}
	// Setup routes
	setupRoutes(app, apiHandler)

	// Start the Fiber app
	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App, apiHandler *api.Handler) {
	app.Post("/tasks", apiHandler.CreateTask)
	app.Put("/tasks/:id/complete", apiHandler.CompleteTask)
	app.Delete("/tasks/:id", apiHandler.DeleteTask)
}
