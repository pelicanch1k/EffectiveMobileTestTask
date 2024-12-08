package handler

import (
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/service"
	"github.com/pelicanch1k/EffectiveMobileTestTask/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
}

func NewHandler(services *service.Service, logger *logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	songs := router.Group("/api/v1")
	{
		songs.GET("/songs", h.getSongs)
		songs.GET("/songs/:id/lyrics", h.getSongLyrics)
		songs.DELETE("/songs/:id", h.deleteSong)

		songs.PUT("/songs", h.updateSong)
	}

	return router
}
