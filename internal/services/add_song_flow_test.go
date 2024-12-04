package services

import (
	"net/http"
	"net/http/httptest"
	"testApp/internal/models"
	"testApp/internal/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Моковый сервер для имитации внешнего сервиса
func mockExternalServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что переданы правильные параметры
		group := r.URL.Query().Get("group")
		song := r.URL.Query().Get("song")

		if group == "Muse" && song == "Supermassive Black Hole" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"releaseDate": "16.07.2006",
				"text": "Ooh baby, don't you know I suffer?",
				"link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
			}`))
			return
		}

		http.Error(w, "Not Found", http.StatusNotFound)
	}))
}

/*

func TestAddSongFlow(t *testing.T) {
	// 1. Поднимаем моковый сервер
	mockServer := mockExternalServer()
	defer mockServer.Close()

	// 2. Подключение к моковой базе данных
	db, mockDB, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := &repositories.SongRepository{DB: sqlxDB}

	// 3. Создаем реальный сервис с замоканным внешним клиентом
	realService := &RealExternalService{Host: mockServer.URL}
	songService := &SongService{
		Repo:           repo,
		ExternalClient: realService,
	}

	// 4. Настраиваем мок для базы данных
	mockDB.ExpectQuery(`INSERT INTO songs`).
		WithArgs(
			"Muse",
			"Supermassive Black Hole",
			"16.07.2006",
			"Ooh baby, don't you know I suffer?",
			"https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		).
		WillReturnRows(sqlmock.NewRows([]string{"id", "group_name", "song_name", "release_date", "text", "link"}).
			AddRow(1, "Muse", "Supermassive Black Hole", "16.07.2006", "Ooh baby, don't you know I suffer?", "https://www.youtube.com/watch?v=Xsp3_a-PMTw"))

	// 5. Входные данные
	inputSong := models.Song{
		Group:    "Muse",
		SongName: "Supermassive Black Hole",
	}

	// 6. Выполняем метод AddSong
	result, err := songService.AddSong(inputSong)

	// 7. Проверяем результат
	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Muse", result.Group)
	assert.Equal(t, "Supermassive Black Hole", result.SongName)
	assert.Equal(t, "16.07.2006", result.ReleaseDate)
	assert.Equal(t, "Ooh baby, don't you know I suffer?", result.Text)
	assert.Equal(t, "https://www.youtube.com/watch?v=Xsp3_a-PMTw", result.Link)

	// 8. Проверяем, что все ожидания базы данных выполнены
	assert.NoError(t, mockDB.ExpectationsWereMet())
}
*/

func TestAddSongFlow(t *testing.T) {
	// 1. Поднимаем моковый сервер
	mockServer := mockExternalServer()
	defer mockServer.Close()

	// 2. Подключение к моковой базе данных
	db, mockDB, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := &repositories.SongRepository{DB: sqlxDB}

	// 3. Создаем реальный сервис с замоканным внешним клиентом
	realService := &RealExternalService{Host: mockServer.URL}
	songService := &SongService{
		Repo:           repo,
		ExternalClient: realService,
	}

	// 4. Настраиваем мок для базы данных
	mockDB.ExpectExec(`INSERT INTO songs`).
		WithArgs(
			"Muse",
			"Supermassive Black Hole",
			"16.07.2006",
			"Ooh baby, don't you know I suffer?",
			"https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Указываем ID и количество затронутых строк

	// 5. Входные данные
	inputSong := models.Song{
		Group:    "Muse",
		SongName: "Supermassive Black Hole",
	}

	// 6. Выполняем метод AddSong
	result, err := songService.AddSong(inputSong)
	require.NoError(t, err, "AddSong вызвал ошибку")

	// 7. Проверяем результат
	require.Equal(t, 1, result.ID)
	require.Equal(t, "Muse", result.Group)
	require.Equal(t, "Supermassive Black Hole", result.SongName)
	require.Equal(t, "16.07.2006", result.ReleaseDate)
	require.Equal(t, "Ooh baby, don't you know I suffer?", result.Text)
	require.Equal(t, "https://www.youtube.com/watch?v=Xsp3_a-PMTw", result.Link)

	// 8. Проверяем, что все ожидания базы данных выполнены
	require.NoError(t, mockDB.ExpectationsWereMet(), "SQLMock ожидания не выполнены")
}
