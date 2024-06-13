package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	redis *redis.Client
	ctx   context.Context
}

type RedisCacheInterface interface {
	SetString(key string, value interface{}, expiration time.Duration) error
	GetString(key string, gettervalue interface{}) (found bool, err error)
	SetHash(key string, value map[string]interface{}, expiration time.Duration) error
	GetHash(key string, field string, gettervalue interface{}) (found bool, err error)
	GetHashAll(key string, gettervalue map[string]interface{}) (found bool, err error)
	Delete(key string) error
	ClearDb() error
	ClearAll() error
	GetMultiple(keys []string, gettervalue []interface{}) (keynotfound []string, err error)
}

func NewRedisCache(redis *redis.Client, ctx context.Context) RedisCacheInterface {
	return &redisCache{
		redis: redis,
		ctx:   ctx,
	}
}

func (r redisCache) SetString(key string, value interface{}, expiration time.Duration) error {
	valuetojson, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.redis.Set(r.ctx, key, valuetojson, expiration).Err()
}

func (r redisCache) GetString(key string, gettervalue interface{}) (found bool, err error) {
	valuetojson, err := r.redis.Get(r.ctx, key).Result()
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal([]byte(valuetojson), gettervalue); err != nil {
		return false, err
	}
	return true, nil
}

func (r redisCache) SetHash(key string, value map[string]interface{}, expiration time.Duration) error {

	for k, v := range value {
		jsonValue, err := json.Marshal(v)
		if err != nil {
			return err
		}
		value[k] = jsonValue
	}

	err := r.redis.HSet(r.ctx, key, value).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r redisCache) GetHash(key string, field string, gettervalue interface{}) (found bool, err error) {
	valuetojson, err := r.redis.HGet(r.ctx, key, field).Result()
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal([]byte(valuetojson), gettervalue); err != nil {
		return false, err
	}
	return true, nil
}

func (r redisCache) GetHashAll(key string, gettervalue map[string]interface{}) (found bool, err error) {
	values, err := r.redis.HGetAll(r.ctx, key).Result()
	if err != nil {
		return false, err
	}

	if len(values) == 0 {
		return false, nil
	}

	for k, v := range values {
		var value interface{}
		if err := json.Unmarshal([]byte(v), &value); err != nil {
			return false, err
		}
		gettervalue[k] = value
	}

	return true, nil
}

func (r redisCache) Delete(key string) error {
	err := r.redis.Del(r.ctx, key).Err()
	if err == redis.Nil {
		return errors.New("key does not exist")
	}

	if err != nil {
		return err
	}

	return nil

}

func (r redisCache) ClearDb() error {
	err := r.redis.FlushDB(r.ctx).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r redisCache) ClearAll() error {
	err := r.redis.FlushAll(r.ctx).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r redisCache) GetMultiple(keys []string, gettervalue []interface{}) (keynotfound []string, err error) {

	if len(keys) == 0 {
		return nil, errors.New("no keys provided")
	}

	valuetojson, err := r.redis.MGet(r.ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	for i, v := range valuetojson {
		if v == nil {
			keynotfound = append(keynotfound, keys[i])
			continue
		}

		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(data, &gettervalue[i]); err != nil {
			return nil, err
		}
	}
	return keynotfound, nil
}
