package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisService struct {
	RedisClient *redis.Client
}

var ctx = context.Background()

// SetRedisString 设置key=>value
func (redisService *RedisService) SetRedisString(key, value string, time time.Duration) error {
	err := redisService.RedisClient.Set(ctx, key, value, time).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetRedisString 获取key
func (redisService *RedisService) GetRedisString(key string) (string, error) {
	val, err := redisService.RedisClient.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return val, err
	}
	return val, nil
}

// HSetRedisString 设置hash
func (redisService *RedisService) HSetRedisString(key, hKey, value string, time time.Duration) error {
	err := redisService.RedisClient.HSet(ctx, key, hKey, value).Err()
	if err != nil {
		return err
	}
	redisService.RedisClient.Expire(ctx, key, time)
	return nil
}

// HGetRedisString 获取hash类型对应的子key的值
func (redisService *RedisService) HGetRedisString(key, hKey string) (string, error) {
	val, err := redisService.RedisClient.HGet(ctx, key, hKey).Result()
	if err != nil && err != redis.Nil {
		return val, err
	}
	return val, nil
}

// SetNx 加锁 return: true-锁成功 false-锁失败
func (redisService *RedisService) SetNx(key string, value interface{}, time time.Duration) bool {
	val := redisService.RedisClient.SetNX(ctx, key, value, time).Val()
	return val
}
