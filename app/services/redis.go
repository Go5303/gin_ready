package services

import (
	"context"
	"gin_ready/global"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisService struct {
}

var RedisService = new(redisService)

var ctx = context.Background()

var rdb *redis.Client

func (*redisService) InitRdb() *redisService {
	rdb = global.App.Redis
	return new(redisService)
}

// SetRedisString 设置key=>value
func (*redisService) SetRedisString(key, value string, time time.Duration) error {
	err := rdb.Set(ctx, key, value, time).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetRedisString 获取key
func (*redisService) GetRedisString(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return val, err
	}
	return val, nil
}

// HSetRedisString 设置hash
func (*redisService) HSetRedisString(key, hKey, value string, time time.Duration) error {
	err := rdb.HSet(ctx, key, hKey, value).Err()
	if err != nil {
		return err
	}
	rdb.Expire(ctx, key, time)
	return nil
}

// HGetRedisString 获取hash类型对应的子key的值
func HGetRedisString(key, hKey string) (string, error) {
	val, err := rdb.HGet(ctx, key, hKey).Result()
	if err != nil && err != redis.Nil {
		return val, err
	}
	return val, nil
}
