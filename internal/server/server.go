package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luthfiupil/dojang/internal/handlers"
)

type Server struct {
	Router *chi.Mux
	DB     *pgxpool.Pool
}

func NewServer(db *pgxpool.Pool) *Server {
	r := chi.NewRouter()

	// Health
	r.Get("/health", handlers.HealthCheck)

	// Roles
	roleHandler := handlers.NewRoleHandler(db)
	r.Get("/roles", roleHandler.GetAllRoles)

	userHandler := handlers.NewUserHandler(db)
	r.Post("/users", userHandler.CreateUser)
	r.Get("/users", userHandler.GetUsers)      // ✅ List users
	r.Get("/users/{id}", userHandler.GetUsers) // Single user (coming soon)

	return &Server{
		Router: r,
		DB:     db,
	}
}
