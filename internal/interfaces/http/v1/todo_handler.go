package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	application "github.com/smile-ko/go-ddd-template/internal/application/todo"
	"github.com/smile-ko/go-ddd-template/pkg/logger"
)

type TodoHandler struct {
	uc  application.ITodoUsecase
	log logger.ILogger
}

func NewTodoHandler(r *gin.Engine, uc application.ITodoUsecase, log logger.ILogger) *TodoHandler {
	return &TodoHandler{
		uc:  uc,
		log: log,
	}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var in application.CreateTodoInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := h.uc.Create(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, out)
}

func (h *TodoHandler) List(c *gin.Context) {
	out, err := h.uc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *TodoHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	out, err := h.uc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}
