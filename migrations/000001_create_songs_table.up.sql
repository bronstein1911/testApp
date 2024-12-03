CREATE TABLE songs (
    id SERIAL PRIMARY KEY,               -- Уникальный идентификатор песни
    group_name VARCHAR(255) NOT NULL,    -- Название группы
    song_title VARCHAR(255) NOT NULL,    -- Название песни
    release_date DATE,                   -- Дата выпуска
    lyrics TEXT,                         -- Текст песни
    link TEXT,                           -- Ссылка на песню
    created_at TIMESTAMP DEFAULT NOW(),  -- Время добавления записи
    updated_at TIMESTAMP DEFAULT NOW()   -- Время последнего обновления записи
);