package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataSources struct {
	DB          *sqlx.DB
	RedisClient *redis.Client
}

func InitDB() (*DataSources, error) {
	log.Printf("Initializing data sources\n")
	envFile, errs := godotenv.Read(".env")

	if errs != nil {
		fmt.Printf("Error reading in env file, %s: %v", envFile["PG_HOST"], errs)
	}
	fmt.Printf(envFile["PG_HOST"])

	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgDB := os.Getenv("PG_DB")

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPass, pgDB)
	fmt.Printf(pgConnString)
	log.Printf("Connecting to db...")
	db, err := sqlx.Open("postgres", pgConnString)

	if err != nil {
		return nil, fmt.Errorf("Error opening db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisDB := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	_, err = redisDB.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("Error connecting to redis: %w", err)
	}

	return &DataSources{
		DB:          db,
		RedisClient: redisDB,
	}, nil
}

func (d *DataSources) Close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}

	if err := d.RedisClient.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}
