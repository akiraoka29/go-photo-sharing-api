package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	router http.Handler
}

func NewServer() *Server {
	// db := database.NewMySQLDB("root@tcp(localhost:3306)/ecommerce")
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, this is a sample REST API!")
	})

	s := &Server{
		router: router,
	}

	return s
}

func (s *Server) Run(addr string) error {
	fmt.Printf("Server started on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}