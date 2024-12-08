package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pelicanch1k/EffectiveMobileTestTask/pkg/logging"
	"github.com/pelicanch1k/EffectiveMobileTestTask/structs"
	"strings"
)

type SongsPostgres struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewSongsPostgres(db *sqlx.DB) *SongsPostgres {
	return &SongsPostgres{db: db, logger: logging.GetLogger()}
}

func (s SongsPostgres) GetSongs(resp structs.GetSongsRequest) ([]structs.Song, error) {
	paramIndex := 1

	// Pagination logic here
	query := "SELECT id, genre, song, releaseDate, text, link FROM songs WHERE 1=1"
	args := []interface{}{}

	if resp.Genre != "" {
		query += fmt.Sprintf(" AND group_name = $%d", paramIndex)
		args = append(args, resp.Genre)
		paramIndex++
	}

	if resp.Song != "" {
		query += fmt.Sprintf(" AND song = $%d", paramIndex)
		args = append(args, resp.Song)
		paramIndex++
	}

	if resp.ReleaseDate != "" {
		query += fmt.Sprintf(" AND releaseDate = $%d", paramIndex)
		args = append(args, resp.ReleaseDate)
		paramIndex++
	}

	if resp.Limit != 0 && resp.Offset != 0 {
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1)
		args = append(args, resp.Limit, resp.Offset)
		paramIndex += 2
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return nil, err
	}
	defer rows.Close()

	var songs []structs.Song

	for rows.Next() {
		var s structs.Song
		if err := rows.Scan(&s.Id, &s.Genre, &s.Song, &s.ReleaseDate, &s.Text, &s.Link); err != nil {
			return nil, err
		}
		songs = append(songs, s)
	}

	return songs, nil
}

func (s SongsPostgres) AddSong(req structs.AddSongRequest) (int, error) {
	query := "INSERT INTO songs (genre, song, releaseDate, text, link) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.db.Exec(query, req.Genre, req.Song, req.ReleaseDate, req.Text, req.Link)
	if err != nil {
		return 0, err
	}

	return req.Id, nil
}

func (s SongsPostgres) UpdateSong(req structs.UpdateSongRequest) error {
	paramIndex := 1

	query := "UPDATE songs SET "
	args := []interface{}{}

	if req.Song != "" {
		query += fmt.Sprintf("song = $%d,", paramIndex)
		args = append(args, req.Song)
		paramIndex++
	}

	if req.Genre != "" {
		query += fmt.Sprintf("genre = $%d,", paramIndex)
		args = append(args, req.Genre)
		paramIndex++
	}

	if req.ReleaseDate != "" {
		query += fmt.Sprintf("releaseDate = $%d,", paramIndex)
		args = append(args, req.ReleaseDate)
		paramIndex++
	}

	if req.Link != "" {
		query += fmt.Sprintf("link = $%d,", paramIndex)
		args = append(args, req.Link)
		paramIndex++
	}

	if req.Text != "" {
		query += fmt.Sprintf("text = $%d,", paramIndex)
		args = append(args, req.Text)
		paramIndex++
	}

	runes := []rune(query)
	if runes[len(runes)-1] == ',' {
		query = string(runes[:len(runes)-1])
	}

	query += fmt.Sprintf(" WHERE id = $%d;", paramIndex)
	args = append(args, req.Id)
	s.logger.Info(query)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (s SongsPostgres) DeleteSong(id int) error {
	query := fmt.Sprintf("DELETE FROM songs WHERE id = $1")

	_, err := s.db.Exec(query, id)
	return err
}

func (s SongsPostgres) GetSongLyrics(req structs.GetSongLyricsRequest) ([]string, error) {

	// Получаем текст песни из базы данных
	var lyrics string

	query := "SELECT text FROM songs WHERE id = $1"

	err := s.db.QueryRow(query, req.Id).Scan(&lyrics)
	if err == sql.ErrNoRows {
		return nil, errors.New("Song not found")
	} else if err != nil {
		return nil, errors.New("Failed to fetch song lyrics")
	}

	// Разделяем текст на куплеты
	verses := splitVerses(lyrics)
	start := (req.Offset - 1) * req.Limit
	end := start + req.Limit
	if start >= len(verses) {
		return verses, nil
	}
	if end > len(verses) {
		end = len(verses)
	}

	return verses[start:end], nil
}

// splitVerses разбивает текст песни на куплеты
func splitVerses(text string) []string {
	return strings.Split(text, "\n\n") // Куплеты разделены двумя переводами строки
}
