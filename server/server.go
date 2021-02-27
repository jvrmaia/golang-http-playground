package server

import (
	"github.com/jvrmaia/golang-http-playground/logger"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	logger *logger.Logger
}

func NewServer(mux ServeMux) (*Server, error) {
	server := Server{}

	server.logger = logger.New(false)

	server.server = &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &server, nil
}

func (s *Server) Start() error {
	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		s.logger.Fatal().Err(err).Msg("server startup failed")
	}

	return err
}
