package services

import (
	"testApp/config"
	"testApp/internal/models"
	"testApp/internal/repositories"
)

type SongService struct {
	Repo           *repositories.SongRepository
	Config         config.Config
	ExternalClient ExternalService
}

func (s *SongService) GetSongs(page, limit int, filters map[string]string) ([]models.Song, error) {
	return s.Repo.GetAllSongs(page, limit, filters)
}

func (s *SongService) AddSong(song models.Song) (models.Song, error) {
	// Запрос во внешний сервис
	externalData, err := s.ExternalClient.FetchSongData(song.Group, song.SongName)
	if err != nil {
		return models.Song{}, err
	}

	// Обогащаем песню данными из внешнего сервиса
	song.ReleaseDate = externalData.ReleaseDate
	song.Text = externalData.Text
	song.Link = externalData.Link

	// Добавляем песню в БД
	addedSong, err := s.Repo.AddSong(song)
	if err != nil {
		return models.Song{}, err
	}

	return addedSong, nil
}

// Другие методы: для PATCH, DELETE
