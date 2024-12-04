package repositories

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"testApp/internal/models"
)

type SongRepository struct {
    DB *sqlx.DB
}

func (r *SongRepository) GetAllSongs(page, limit int, filters map[string]string) ([]models.Song, error) {
    query := "SELECT id, group, song, releaseDate, text, link FROM songs WHERE 1=1"
    
    if group, ok := filters["group"]; ok {
        query += fmt.Sprintf(" AND group = '%s'", group)
    }
    if song, ok := filters["song"]; ok {
        query += fmt.Sprintf(" AND song LIKE '%%%s%%'", song)
    }

    query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, (page-1)*limit)
    
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var songs []models.Song
    for rows.Next() {
        var song models.Song
        if err := rows.Scan(&song.ID, &song.Group, &song.SongName, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
            return nil, err
        }
        songs = append(songs, song)
    }
    return songs, nil
}

// AddSong добавляет новую песню в базу данных
func (r *SongRepository) AddSong(song models.Song) (models.Song, error) {
	query := `
		INSERT INTO songs (group_name, song_name, release_date, text, link)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, group_name, song_name, release_date, text, link
	`

	// Используем sqlx для выполнения запроса и привязки результата к структуре
	var savedSong models.Song
	err := r.DB.QueryRowx(
		query,
		song.Group,
		song.SongName,
		song.ReleaseDate,
		song.Text,
		song.Link,
	).StructScan(&savedSong)

	if err != nil {
		return models.Song{}, fmt.Errorf("failed to add song to database: %w", err)
	}

	return savedSong, nil
}