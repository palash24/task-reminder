package internal

import (
	"github.com/go-chi/chi"
	"github.com/palash24/task-reminder/config"
	api "github.com/palash24/task-reminder/internal/api"
)

type Server struct {
	Router     *chi.Mux
	ServerAddr string
	ServerPort string
}

func NewHttpServer() Server {
	cfg := config.NewConfig()

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Post("/api/v1/tasks", api.CreateNewTask)
		r.Get("/api/v1/tasks", api.GetAllTheTasks)
		r.Get("/api/v1/tasks/{taskID}", api.GetIndividualTasks)
		r.Put("/api/v1/tasks/{taskID}", api.UpdateTaskByTaskID)
		r.Delete("/api/v1/tasks/{taskID}", api.DeleteTaskByTaskID)

		r.Post("/api/v1/reminders", api.CreateNewReminder)
		r.Get("/api/v1/reminders", api.GetAllReminders)
		r.Get("/api/v1/reminders/{reminderID}", api.GetSingleReminderByID)
		r.Put("/api/v1/reminders/{reminderID}", api.UpdateReminderByID)
		r.Delete("/api/v1/reminders/{reminderID}", api.DeleteReminderByID)
	})
	return Server{
		Router:     router,
		ServerAddr: cfg.ServerAddr,
		ServerPort: cfg.ServerPort,
	}
}
