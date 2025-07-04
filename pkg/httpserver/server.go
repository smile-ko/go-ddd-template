// Package httpserver implements HTTP Server.
package httpserver

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smile-ko/go-ddd-template/config"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	Router          *gin.Engine
}

func New(cfg *config.Config, router *gin.Engine) *Server {
	server := prepareHttpServer(cfg, router)

	return server
}

func prepareHttpServer(cfg *config.Config, router *gin.Engine) *Server {
	httpServer := &http.Server{
		Handler:      router,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}
	httpServer.Addr = net.JoinHostPort("", cfg.HTTP.Port)

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		Router:          router,
	}
	return s
}

func (s *Server) Start() {
	go func() {
		println("[HTTP] Server is running at 0.0.0.0" + s.server.Addr)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
