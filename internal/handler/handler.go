package handler

import (
	"github.com/gin-gonic/gin"
	"scheduler-server/internal/database"
)

type Handler struct {
	db database.IDB
}

func NewHandler(db database.IDB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "my health is excelente",
	})
}
