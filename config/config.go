package config

type Config struct {
	Environment         string `env:"ENVIRONMENT,required"`
	SongInfoServiceHost string `env:"SONG_INFO_SERVICE_HOST,required"`
	MigrationsPath      string `env:"MIGRATION_PATH,required"`
	PostgresUrl         string `env:"POSTGRES_URL,required"`
}
