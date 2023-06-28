package http

import (
	"log"
	"net/http"
)

// Server represents the HTTP server.
type Server struct {
	Addr    string
	Handler http.Handler
}

// NewServer creates a new instance of the HTTP server.
func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		Addr:    addr,
		Handler: handler,
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	log.Printf("Server listening on %s\n", s.Addr)
	return http.ListenAndServe(s.Addr, s.Handler)
}
