package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testApp/internal/models"
)

type RealExternalService struct {
	Host string // Хост внешнего сервиса
}

func (s *RealExternalService) FetchSongData(group, song string) (models.Song, error) {
	// Формируем URL для запроса
	url := fmt.Sprintf("%s/info?group=%s&song=%s", s.Host, group, song)

	resp, err := http.Get(url)
	if err != nil {
		return models.Song{}, fmt.Errorf("failed to fetch external song data: %w", err)
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var externalData struct {
		ReleaseDate string `json:"releaseDate"`
		Text        string `json:"text"`
		Link        string `json:"link"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&externalData); err != nil {
		return models.Song{}, fmt.Errorf("failed to decode external response: %w", err)
	}

	return models.Song{
		ReleaseDate: externalData.ReleaseDate,
		Text:        externalData.Text,
		Link:        externalData.Link,
	}, nil
}
