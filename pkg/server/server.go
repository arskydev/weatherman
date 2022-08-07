package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	SHUTDOWN_TIMEOUT = 5 * time.Second
)

type Server struct {
	httpServer *http.Server
	handler    http.Handler
}

func NewServer(h http.Handler) *Server {
	return &Server{
		handler: h,
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Run(ctx context.Context, port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        s.handler,
		MaxHeaderBytes: 1 << 20, // 1048576 byte -> 1024 kbyte -> 1 Mbyte
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			switch err.Error() {
			case "http: Server closed":
				log.Println(err.Error())
			default:
				log.Fatalf("Error raised while server Listen And Serve:\n%s", err.Error())
			}
		}
	}()

	log.Printf("Server started on port %v\n", port)
	<-ctx.Done()

	log.Println("Shutting down server gracefully...")

	if err := s.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error raised while server shutdown:\n%w", err)
	}

	log.Println("Shutting down server gracefully... Done")

	return nil
}
