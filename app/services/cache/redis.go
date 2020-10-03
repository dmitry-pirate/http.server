package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/basketforcode/http.server/config"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Redis struct {
	config     *config.Config
	connection *redis.Client
}

//new connection to redis store
func New(config *config.Config) Redis {
	db, _ := strconv.Atoi(config.Redis.DBIndex)
	return Redis{
		config: config,
		connection: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.Redis.DBHost, config.Redis.DBPort),
			Password: config.Redis.DBPassword,
			DB:       db,
		}),
	}
}

//set value to redis
func (r *Redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	model, err := json.Marshal(value)
	err = r.connection.Set(ctx, key, model, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

//get value from redis by key string
func (r *Redis) Get(ctx context.Context, key string) (interface{}, error) {
	valueString, err := r.connection.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, err
	} else {
		var model interface{}
		err = json.Unmarshal([]byte(valueString), &model)
		if err != nil {
			return nil, err
		}
		return model, nil
	}
}
