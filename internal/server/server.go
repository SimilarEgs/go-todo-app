package server

import (
	"context"
	"net/http"
	"time"
)

// use this stuct to create custom server
type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(port string, handler http.Handler) error {

	s.httpServer = &http.Server{
		// adjusting server settings
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// use when exiting application
func (s *Server) ShutDownServer(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
