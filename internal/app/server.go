package app

import (
	"log"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/handlers"
	"social-network/internal/repository"
)

type Server struct {
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
}

func NewServer(db *repository.Database, cfg *config.Config) *Server {
	userRepo := repository.NewUserRepository(db.DB, cfg)
	return &Server{
		authHandler: handlers.NewAuthHandler(userRepo, cfg),
		userHandler: handlers.NewUserHandler(userRepo, cfg),
	}
}

func (s *Server) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /login", s.authHandler.Login)
	mux.HandleFunc("POST /user/register", s.userHandler.Register)
	mux.HandleFunc("GET /user/get/{id}", s.userHandler.GetUser)
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()
	s.SetupRoutes(mux)

	log.Printf("Server starting on %s", addr)
	return http.ListenAndServe(addr, mux)
}
