package server

import (
	"fmt"
	"net/http"

	"github.com/akiraoka29/go-photo-sharing-api/api/handlers"
	"github.com/akiraoka29/go-photo-sharing-api/api/middleware"
	"github.com/akiraoka29/go-photo-sharing-api/pkg/database"
)

type Server struct {
	router http.Handler
}

func NewServer() *Server {
	db := database.NewMySQLDB("root@tcp(localhost:3306)/ecommerce")
	userHandler := handlers.NewUserHandler(db)
	authMiddleware := middleware.NewAuthMiddleware()

	router := http.NewServeMux()
	router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/register", userHandler.Register)

	s := &Server{
		router: router,
	}

	return s
}

func (s *Server) Run(addr string) error {
	fmt.Printf("Server started on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}