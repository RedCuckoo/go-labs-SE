package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	*sql.DB
}

func Configure() (Config, error) {
	var databaseUrl struct{
		DatabaseUrl string `split_words:"true" required:"true"`
	}

	err := envconfig.Process("POSTGRES", &databaseUrl)

	if err != nil {
		log.Fatalf("failed to get environment variable, err: %v", err)
		return Config{}, err
	}

	database, err := sql.Open("postgres", databaseUrl.DatabaseUrl)

	if err != nil {
		return Config{}, err
	}

	return Config{DB: database}, nil
}
