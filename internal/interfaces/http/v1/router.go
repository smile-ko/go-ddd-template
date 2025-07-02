package http

import (
	"github.com/gin-gonic/gin"
	application "github.com/smile-ko/go-ddd-template/internal/application/todo"
	"github.com/smile-ko/go-ddd-template/pkg/logger"
)

// @title Example API v1
// @version 1.0
// @description This is the v1 of the API
// @BasePath /api/v1
func NewRouterV1(r *gin.Engine, todoUsecase application.ITodoUsecase, log logger.ILogger) {
	todoHandler := NewTodoHandler(r, todoUsecase, log)

	v1 := r.Group("/api/v1")
	// Define the routes for the todo resource
	{
		v1.POST("/todos", todoHandler.Create)
		v1.GET("/todos", todoHandler.List)
		v1.GET("/todos/:id", todoHandler.Get)
	}
	// Add more routes as needed
}
