package models

type Song struct {
	ID          int    `json:"id"`
	Group       string `json:"group"`
	SongName    string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
