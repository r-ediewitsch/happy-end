package main

import (
	"log"
	"net/http"

	"happy-end/internal/config"
	"happy-end/internal/database"
	"happy-end/internal/handler"
	"happy-end/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	logsRepo := repository.NewLogsRepository(db)
	logsHandler := handler.NewLogHandler(logsRepo)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":   "unhealthy",
				"database": "disconnected",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   "healthy",
			"database": "connected",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/logs", logsHandler.GetLogs)
		api.GET("/logs/:id", logsHandler.GetLogById)
	}

	if cfg.AppUrl == "development" {
		dev := r.Group("/dev")
		{
			dev.POST("/logs", logsHandler.CreateLog)
			dev.PUT("/logs/:id", logsHandler.UpdateLog)
			dev.DELETE("/logs/:id", logsHandler.DeleteLog)
		}
	}

	log.Printf("Server listening on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
