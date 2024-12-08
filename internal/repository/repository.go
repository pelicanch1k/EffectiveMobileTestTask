package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/repository/postgres"
	"github.com/pelicanch1k/EffectiveMobileTestTask/structs"
)

type Songs interface {
	GetSongs(resp structs.GetSongsRequest) ([]structs.Song, error)
	AddSong(song structs.AddSongRequest) (int, error)
	UpdateSong(song structs.UpdateSongRequest) error
	DeleteSong(id int) error
	GetSongLyrics(req structs.GetSongLyricsRequest) ([]string, error)
}

type Repository struct {
	Songs
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Songs: postgres.NewSongsPostgres(db),
	}
}
