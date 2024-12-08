package structs

type Song struct {
	Id int `json:"id"`

	Genre string `json:"genre"`
	Song  string `json:"song"`

	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type UpdateSongRequest struct {
	Id int `json:"id" binding:"required"`

	Genre string `json:"genre"`
	Song  string `json:"song"`

	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type AddSongRequest struct {
	Id int `json:"id"`

	Genre string `json:"genre" binding:"required"`
	Song  string `json:"song" binding:"required"`

	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type GetSongsRequest struct {
	Genre       string
	Song        string
	ReleaseDate string

	Limit  int
	Offset int
}

type GetSongLyricsRequest struct {
	Id int

	Limit  int
	Offset int
}
