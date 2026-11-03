package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server contains the application's logger and configured HTTP server.
type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

// New creates a server with all application routes registered.
func New(logger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/upload", handlers.Upload)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{Logger: logger, HTTPServer: httpServer}
}

// ListenAndServe starts the configured HTTP server.
func (s *Server) ListenAndServe() error {
	return s.HTTPServer.ListenAndServe()
}
