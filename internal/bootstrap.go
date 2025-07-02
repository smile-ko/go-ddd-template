package internal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/smile-ko/go-ddd-template/config"
	application "github.com/smile-ko/go-ddd-template/internal/application/todo"
	"github.com/smile-ko/go-ddd-template/internal/infrastructure/db/sqlc"
	"github.com/smile-ko/go-ddd-template/internal/infrastructure/repository"
	v1 "github.com/smile-ko/go-ddd-template/internal/interfaces/http/v1"
	"github.com/smile-ko/go-ddd-template/pkg/httpserver"
	"github.com/smile-ko/go-ddd-template/pkg/logger"
	"github.com/smile-ko/go-ddd-template/pkg/postgres"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	// Logger setup
	log := logger.NewLogger(cfg)

	// DB connection
	pg := postgres.NewOrGetSingleton(cfg)
	defer pg.Close()

	// HTTP router setup
	var router *gin.Engine
	if cfg.App.EnvName == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		// Add logging & recovery for production
		router.Use(gin.Recovery())
	}

	// Initialize database queries
	queries := sqlc.New(pg.Pool)

	// Todo repository and use case
	todoRepo := repository.NewTodoRepository(queries)
	todoUseCase := application.NewTodoUseCase(todoRepo)

	// Register routes
	v1.NewTodoHandler(router, todoUseCase, log)

	// HTTP server
	httpServer := httpserver.New(cfg, router)

	// Start server
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
	select {
	case sig := <-interrupt:
		log.Info("app - Run - signal received", zap.String("signal", sig.String()))

	case err = <-httpServer.Notify():
		log.Error("app - Run - httpServer.Notify error", zap.Error(err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app - Run - httpServer.Shutdown error", zap.Error(err))
	}
}
