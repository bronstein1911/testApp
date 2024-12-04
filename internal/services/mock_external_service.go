package services

import "testApp/internal/models"

// MockExternalService - моковая реализация внешнего сервиса
type MockExternalService struct {
    MockResponse models.Song
    MockError    error
}

func (m *MockExternalService) FetchSongData(group, song string) (models.Song, error) {
    // Возвращаем заранее заданный ответ или ошибку
    return m.MockResponse, m.MockError
}
