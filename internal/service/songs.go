package service

import (
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/repository"
	"github.com/pelicanch1k/EffectiveMobileTestTask/structs"
)

type SongsService struct {
	repo *repository.Repository
}

func NewSongsService(repo *repository.Repository) *SongsService {
	return &SongsService{repo: repo}
}

func (s SongsService) GetSongs(req structs.GetSongsRequest) ([]structs.Song, error) {
	return s.repo.GetSongs(req)
}

func (s SongsService) AddSong(req structs.AddSongRequest) (int, error) {
	song, err := getRequest(req)
	if err != nil {
		return 0, err
	}

	return s.repo.AddSong(*song)
}

func (s SongsService) UpdateSong(req structs.UpdateSongRequest) error {
	return s.repo.UpdateSong(req)
}

func (s SongsService) DeleteSong(id int) error {
	return s.repo.DeleteSong(id)
}

func (s SongsService) GetSongLyrics(req structs.GetSongLyricsRequest) ([]string, error) {
	return s.repo.GetSongLyrics(req)
}
