package service

import (
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/repository"
	"github.com/pelicanch1k/EffectiveMobileTestTask/structs"
)

type Songs interface {
	GetSongs(resp structs.GetSongsRequest) ([]structs.Song, error)
	AddSong(song structs.AddSongRequest) (int, error)
	UpdateSong(song structs.UpdateSongRequest) error
	DeleteSong(id int) error
	GetSongLyrics(req structs.GetSongLyricsRequest) ([]string, error)
}

type Service struct {
	Songs
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Songs: NewSongsService(repo),
	}
}
