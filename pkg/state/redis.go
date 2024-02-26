package state

import "github.com/redis/go-redis/v9"

func NewRedisClient(url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)

	if err != nil {
		return &redis.Client{}, err
	}

	return redis.NewClient(opts), nil
}