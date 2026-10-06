package handler

import (
	"net/http"
	"strconv"

	"happy-end/internal/model"
	"happy-end/internal/repository"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	repo *repository.LogsRepository
}

func NewLogHandler(repo *repository.LogsRepository) *LogHandler {
	return &LogHandler{repo: repo}
}

func (h *LogHandler) GetLogs(c *gin.Context) {
	logs, err := h.repo.GetLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func (h *LogHandler) GetLogById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	log, err := h.repo.GetLogById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

func (h *LogHandler) CreateLog(c *gin.Context) {
	var log model.Logs
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdLog, err := h.repo.CreateLog(&log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdLog)
}

func (h *LogHandler) UpdateLog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var log model.Logs
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedLog, err := h.repo.UpdateLog(id, &log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedLog)
}

func (h *LogHandler) DeleteLog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.repo.DeleteLog(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Log deleted successfully"})
}
