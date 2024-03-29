package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/go-redis/redis/v8"
)

type redisRepo struct {
	Redis *redis.Client
}

func NewTokenRepository(redisClient *redis.Client) model.TokenRepository {
	return &redisRepo{
		Redis: redisClient,
	}
}

func (r *redisRepo) SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error {
	key := fmt.Sprintf("%s:%s", userID, tokenID)

	if err := r.Redis.Set(ctx, key, 0, expiresIn).Err(); err != nil {
		log.Printf("Could not SET refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return model.NewInternal()
	}

	return nil
}

func (r *redisRepo) DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error {
	key := fmt.Sprintf("%s:%s", userID, prevTokenID)

	result := r.Redis.Del(ctx, key)

	if err := result.Err(); err != nil {
		log.Printf("Could not delete refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, prevTokenID, err)
		return model.NewInternal()
	}

	if result.Val() < 1 {
		log.Printf("Refresh token to redis for userID/tokenID: %s/%s does not exist\n", userID, prevTokenID)
		return model.NewAuthorization("Unable to authenticate user")
	}

	return nil
}

func (r *redisRepo) DeleteUserRefreshToken(ctx context.Context, userID string) error {
	pattern := fmt.Sprintf("%s*", userID)

	// Scan for this pattern
	iter := r.Redis.Scan(ctx, 0, pattern, 5).Iterator()
	failCount := 0

	for iter.Next(ctx) {
		if err := r.Redis.Del(ctx, iter.Val()).Err(); err != nil {
			log.Printf("Failed to delete refresh Token: %s\n", iter.Val())
			failCount++
		}
	}

	if err := iter.Err(); err != nil {
		log.Printf("Failed to delete refresh token: %s\n", iter.Val())
	}

	if failCount > 0 {
		return model.NewInternal()
	}

	return nil
}
