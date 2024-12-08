package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pelicanch1k/EffectiveMobileTestTask/structs"
	"net/http"
	"os"
)

func getRequest(req structs.AddSongRequest) (*structs.AddSongRequest, error) {
	var songDetail *structs.AddSongRequest

	resp, err := http.Get(fmt.Sprintf(os.Getenv("URL_ADD_SONG")+"/info?group=%s&song=%s", req.Genre, req.Song))
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to fetch external data")
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, errors.New("Failed to decode external response")
	}

	return songDetail, nil
}
