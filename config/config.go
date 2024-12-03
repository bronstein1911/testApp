package config

type Config struct {
	Environment    string `env:"ENVIRONMENT,required"`
	MigrationsPath string `env:"MIGRATION_PATH,required"`
	PostgresUrl    string `env:"POSTGRES_URL,required"`
}
