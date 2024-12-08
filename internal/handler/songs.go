package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pelicanch1k/EffectiveMobileTestTask/structs"
	"net/http"
	"strconv"
)

type pagination struct {
	limit, offset int
}

func initPagination(c *gin.Context) (*pagination, error) {
	limit, err := strconv.Atoi(c.GetHeader("limit"))
	if err != nil {
		if c.GetHeader("limit") != "" {
			newErrorResponse(c, http.StatusBadRequest, "Invalid limit parameter")
			return nil, err
		}
	}

	offset, err := strconv.Atoi(c.GetHeader("offset"))
	if err != nil {
		if c.GetHeader("offset") != "" {
			newErrorResponse(c, http.StatusBadRequest, "Invalid offset parameter")
			return nil, err
		}
	}

	return &pagination{limit: limit, offset: offset}, nil
}

func (h *Handler) getSongs(c *gin.Context) {
	genre := c.GetHeader("genre")
	song := c.GetHeader("song")
	releaseDate := c.GetHeader("releaseDate")

	pag, err := initPagination(c)
	if err != nil {
		return
	}

	resp := structs.GetSongsRequest{
		genre, song, releaseDate, pag.limit, pag.offset,
	}

	songs, err := h.services.GetSongs(resp)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *Handler) getSongLyrics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid Song ID")
		return
	}

	pag, err := initPagination(c)
	if err != nil {
		return
	}

	resp := structs.GetSongLyricsRequest{
		id, pag.limit, pag.offset,
	}

	verses, err := h.services.GetSongLyrics(resp)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"verses": verses,
	})
}

func (h *Handler) deleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid Song ID")
		return
	}

	if err = h.services.DeleteSong(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) updateSong(c *gin.Context) {
	var song structs.UpdateSongRequest

	if err := c.BindJSON(&song); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.UpdateSong(song); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})

}

func (h *Handler) addSong(c *gin.Context) {
	var song structs.AddSongRequest

	if err := c.BindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	id, err := h.services.AddSong(song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Song added successfully"})
}
