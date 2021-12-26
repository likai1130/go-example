package gredis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go-example/common/logger"
	"go-example/config"
)

var redisClient redis.Cmdable

func NewRedisClient() (client redis.Cmdable, err error) {
	if redisClient != nil {
		return redisClient, nil
	}

	redisConfig := config.AppConfig.RedisConfig
	if redisConfig.IsSentinel == false && len(redisConfig.Addrs) == 0 {
		return nil, errors.New("redis addr is absent")
	}

	if len(redisConfig.Addrs) > 1 {
		if client, err = newClusterRedis(redisConfig); err != nil {
			return nil, err
		}
		return client, nil
	}

	if redisConfig.IsSentinel && len(redisConfig.SentinelAddrs) > 1 {
		if client, err = newSentinelRedis(redisConfig); err != nil {
			return nil, err
		}
		return client, nil
	}

	if client, err = newSingleRedis(redisConfig); err != nil {
		return nil, err
	}
	return client, nil

}

func newSingleRedis(c config.RedisConfig) (redis.Cmdable, error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.Addrs[0],
		Password: c.Pwd,
		PoolSize: c.PoolSize,
		DB:       c.DB,
	})

	if err := redisClient.Ping().Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("redis connection failed: %s", err.Error()))
	}
	logger.GetLogger().Info("single redis client init success! ")
	return redisClient, nil
}

func newClusterRedis(c config.RedisConfig) (redis.Cmdable, error) {
	redisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    c.Addrs,
		Password: c.Pwd,
		PoolSize: c.PoolSize,
	})

	if err := redisClient.Ping().Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("redis connection failed: %s", err.Error()))
	}
	logger.GetLogger().Info("cluster redis client init success! ")
	return redisClient, nil
}

func newSentinelRedis(c config.RedisConfig) (redis.Cmdable, error) {
	redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    c.MasterName,
		SentinelAddrs: c.SentinelAddrs,
		Password:      c.Pwd,
		PoolSize:      c.PoolSize,
	})

	if err := redisClient.Ping().Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("redis connection failed: %s", err.Error()))
	}
	logger.GetLogger().Info("sentinel redis client init success! ")
	return redisClient, nil
}

func SetUp() {
	if _, err := NewRedisClient(); err != nil {
		panic(err)
	}
}
