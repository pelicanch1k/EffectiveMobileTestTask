package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/service"
	"github.com/pelicanch1k/EffectiveMobileTestTask/pkg/logging"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/pelicanch1k/EffectiveMobileTestTask/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	songs := router.Group("/api/v1")
	{
		songs.GET("/songs", h.getSongs)
		songs.GET("/song/:id/lyrics", h.getSongLyrics)

		songs.DELETE("/song/:id", h.deleteSong)

		songs.PUT("/song", h.updateSong)
		songs.POST("/song")
	}

	return router
}
