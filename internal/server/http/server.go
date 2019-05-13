package http

import (
	"context"
	"github.com/evzpav/documents-crud-refactored/domain"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type Server struct {
	server   *http.Server
	debugLog logger
	errorLog logger
}

func New(port string, documentService domain.DocumentService, debugLog logger, errorLog logger) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      NewHandler(documentService, debugLog, errorLog),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
		debugLog: debugLog,
		errorLog: errorLog,
	}
}

func (s *Server) ListenAndServe() {
	//go func() {
		log.Print("inside listen and serve")
		s.debugLog.Printf("Documents CRUD running on %s!", s.server.Addr)

		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.errorLog.Printf("error on ListenAndServe: %q", err)
		}
	//}()
}

func (s *Server) Shutdown() {
	s.debugLog.Printf("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		s.errorLog.Printf("could not shutdown in 60s: %q", err)
		return
	}

	s.debugLog.Printf("server gracefully stopped")
}
