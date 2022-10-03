package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataSources struct {
	DB *sqlx.DB
}

func InitDB() (*DataSources, error) {
	log.Printf("Initializing data sources\n")
	envFile, _ := godotenv.Read(".env")

	pgHost := envFile["PG_HOST"]
	pgPort := envFile["PG_PORT"]
	pgUser := envFile["PG_USER"]
	pgPass := envFile["PG_PASS"]
	pgDB := envFile["PG_DB"]

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPass, pgDB)

	log.Printf("Connecting to db...")
	db, err := sqlx.Open("postgres", pgConnString)

	if err != nil {
		return nil, fmt.Errorf("Error opening db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return &DataSources{
		DB: db,
	}, nil
}

func (d *DataSources) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}

	return nil
}
