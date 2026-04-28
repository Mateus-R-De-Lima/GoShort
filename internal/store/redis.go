package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface {
	Get(ctx context.Context, code string) (string, error)
	SaveShortenedURL(ctx context.Context, url string) (string, error)
}

func NewStore(rdb *redis.Client) Store {
	return store{rdb}
}

func (s store) Get(ctx context.Context, code string) (string, error) {
	str, err := s.rdb.HGet(ctx, "encurtador", code).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", fmt.Errorf("failed to get code from encurtador hashmap: %w", err)
	}

	return str, nil
}

func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {

	var code string

	for range 5 {
		code = genCode()

		if err := s.rdb.HGet(ctx, "encurtador", code).Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}
			return "", fmt.Errorf("failed to get code from encurtador hashmap: %w", err)
		}
	}

	if err := s.rdb.HSet(ctx, "encurtador", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to save code and url in encurtador hashmap: %w", err)
	}

	return code, nil
}
