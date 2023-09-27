package main

import (
	"log"

	"event-todo/pkg/api"
	"event-todo/pkg/events"
	"event-todo/pkg/repo"
	"event-todo/pkg/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	eventStore := repo.NewInMemoryEventStore()


	projectionManager := events.NewProjectionManager()

	commandHandler := &todo.CommandHandler{
		EventStore: 	  eventStore,
		ProjectionManager: projectionManager,
	}

	apiHandler	:= &api.Handler{
		CommandHandler: commandHandler,
		ProjectionManager: projectionManager,
	}
	setupRoutes(app, apiHandler)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App, apiHandler *api.Handler) {
	app.Post("/tasks", apiHandler.CreateTask)
	app.Put("/tasks/:id/complete", apiHandler.CompleteTask)
	app.Delete("/tasks/:id", apiHandler.DeleteTask)
}
