package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/basketforcode/http.server/pkg/config"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
	"time"
)

type Redis struct {
	config     *config.Config
	connection *redis.Client
}

func New(config *config.Config) Redis {
	return Redis{
		config:     config,
		connection: getClient(config),
	}
}

func getClient(conf *config.Config) *redis.Client {
	db, _ := strconv.Atoi(conf.Redis.DBIndex)

	if conf.Redis.Driver == config.CacheDriverRedisSentinel {
		return redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    conf.Redis.SentinelDBService,
			SentinelAddrs: strings.Split(conf.Redis.SentinelDBHosts, ","),
			Password:      conf.Redis.SentinelDBPassword,
			DB:            db,
		})
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.DBHost, conf.Redis.DBPort),
		Password: conf.Redis.DBPassword,
		DB:       db,
	})
}

func (r *Redis) Close() error {
	return r.connection.Close()
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	model, err := json.Marshal(value)
	err = r.connection.Set(ctx, key, model, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (interface{}, error) {
	valueString, err := r.connection.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, err
	}

	var model interface{}

	err = json.Unmarshal([]byte(valueString), &model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Redis) Ping(ctx context.Context) *redis.StatusCmd {
	return r.connection.Ping(ctx)
}

func (r *Redis) Unlink(ctx context.Context, key string) error {
	err := r.connection.Unlink(ctx, key).Err()

	if err != nil {
		return err
	}

	return nil
}
