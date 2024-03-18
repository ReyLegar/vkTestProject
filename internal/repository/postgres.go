package repository

import (
	"database/sql"
	"log"

	"github.com/ReyLegar/vkTestProject/internal/config"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)

	if err != nil {
		log.Fatal("No connect data base")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("No ping data base")
		return nil, err
	}

	return db, nil
}
