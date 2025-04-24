package redis

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheManager struct {
	client *redis.Client
	ttl    time.Duration
}

// NewCacheManager creates a new instance of CacheManager.
func NewCacheManager(client *redis.Client, ttl time.Duration) *CacheManager {
	return &CacheManager{
		client: client,
		ttl:    ttl,
	}
}

// Get retrieves a value from the cache and deserializes it.
func (cm *CacheManager) Get(ctx context.Context, key string, dest interface{}) (bool, error) {
	cached, err := cm.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// Key not found in cache
		return false, nil
	}
	if err != nil {
		return false, err
	}

	// Deserialize the cached value
	if err := json.Unmarshal([]byte(cached), dest); err != nil {
		return false, err
	}

	log.Printf("Cache hit for key: %s", key)
	return true, nil
}

// Set serializes a value and stores it in the cache.
func (cm *CacheManager) Set(ctx context.Context, key string, value interface{}) error {
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return cm.client.Set(ctx, key, serialized, cm.ttl).Err()
}

// Delete removes one or more keys from the cache.
func (cm *CacheManager) Delete(ctx context.Context, keys ...string) error {
	return cm.client.Del(ctx, keys...).Err()
}
