package services

import "testApp/internal/models"

type ExternalService interface {
    FetchSongData(group, song string) (models.Song, error)
}