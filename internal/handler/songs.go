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

// @Summary Get songs
// @Tags songs
// @Description Get a list of songs with optional filters and pagination
// @ID getSongs
// @Accept  json
// @Produce  json
// @Param  genre header string false "Filter by genre"
// @Param  song header string false "Filter by song name"
// @Param  releaseDate header string false "Filter by release date"
// @Param  limit query int false "Pagination limit"
// @Param  offset query int false "Pagination offset"
// @Success 200 {array} structs.Song
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /songs [get]
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

// @Summary Get song lyrics
// @Security ApiKeyAuth
// @Tags songs
// @Description Get lyrics of a song by its ID with optional pagination
// @ID getSongLyrics
// @Accept  json
// @Produce  json
// @Param  id path int true "Song ID"
// @Param  limit header int false "Pagination limit"
// @Param  offset header int false "Pagination offset"
// @Success 200 {object} map[string][]string "Lyrics of the song"
// @Failure 400 {object} errorResponse "Invalid Song ID"
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /songs/{id}/lyrics [get]
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

// @Summary Delete a song
// @Security ApiKeyAuth
// @Tags songs
// @Description Delete a song by its ID
// @ID deleteSong
// @Accept  json
// @Produce  json
// @Param  id path int true "Song ID"
// @Success 200 {object} statusResponse "Status OK"
// @Failure 400 {object} errorResponse "Invalid Song ID"
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /songs/{id} [delete]
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

// @Summary Update a song
// @Security ApiKeyAuth
// @Tags songs
// @Description Update an existing song
// @ID updateSong
// @Accept  json
// @Produce  json
// @Param  song body structs.UpdateSongRequest true "Song details to update"
// @Success 200 {object} statusResponse "Status OK"
// @Failure 400 {object} errorResponse "Invalid input data"
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /songs [put]
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

// @Summary Add a new song
// @Security ApiKeyAuth
// @Tags songs
// @Description Add a new song to the catalog
// @ID addSong
// @Accept  json
// @Produce  json
// @Param  song body structs.AddSongRequest true "Details of the song to add"
// @Success 201 {object} map[string]interface{} "Song created successfully"
// @Failure 400 {object} errorResponse "Invalid JSON"
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /songs [post]
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
