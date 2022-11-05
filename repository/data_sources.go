package repository

import (
	"context"
	"fmt"
	"log"

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
	envFile, errs := godotenv.Read("app.env")

	if errs != nil {
		fmt.Printf("Error reading in env file, %s: %v", envFile["PG_HOST"], errs)
	}

	pgHost := envFile["PG_HOST"]
	pgPort := envFile["PG_PORT"]
	pgUser := envFile["PG_USER"]
	pgPass := envFile["PG_PASS"]
	pgDB := envFile["PG_DB"]

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPass, pgDB)

	log.Printf("Connecting to db...")
	fmt.Printf(pgConnString)

	db, err := sqlx.Open("postgres", pgConnString)

	if err != nil {
		return nil, fmt.Errorf("Error opening db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	redisHost := envFile["REDIS_HOST"]
	redisPort := envFile["REDIS_PORT"]

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
