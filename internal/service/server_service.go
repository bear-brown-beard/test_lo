package service

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test_lo/internal/handler"
	"test_lo/internal/logger"
	"time"
)

type ServerService struct {
	logger *logger.Logger
	routes *handler.Routes
	server *http.Server
}

func BuildServerService(logger *logger.Logger, routes *handler.Routes) *ServerService {
	return &ServerService{
		logger: logger,
		routes: routes,
		server: &http.Server{
			Addr: ":8080",
		},
	}
}

func (s *ServerService) Start() {
	s.server.Handler = s.routes.SetupRoutes()
	s.logger.Log("INFO", "Server started on port 8080")

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Log("FATAL", "Could not start server: "+err.Error())
		}
	}()
}

func (s *ServerService) HandleStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Log("ERROR", "Server shutdown failed: "+err.Error())
	}
}
