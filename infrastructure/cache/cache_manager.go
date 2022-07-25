package cache

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"akawork.io/infrastructure/logger"
	"akawork.io/infrastructure/util"
	"github.com/go-redis/redis"
)

/**
 * Defines a CacheManager
 */
type CacheManager struct {
	Client *redis.Client
}

/**
 * Initializes cache
 */
func (manager *CacheManager) Init(host string, poolSize int, minIdleConns int, DB int) {

	manager.Client = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     "", // no password set
		DB:           DB,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
	})
}

// Pipeline
func (manager *CacheManager) Pipeline() redis.Pipeliner {
	return manager.Client.Pipeline()
}

/**
 * Deletes item
 */
func (manager *CacheManager) HIncrBy(key string, field string, incr int64) int64 {
	err := manager.Client.HIncrBy(key, field, incr).Val()
	return err
}

/**
 * Deletes item
 */
func (manager *CacheManager) DeleteItem(key string) int64 {
	err := manager.Client.Del(key).Val()
	return err
}

/**
 * Pushes item to queue list
 */
func (manager *CacheManager) PushItem(queueKey string, value interface{}) error {
	err := manager.Client.RPush(queueKey, value).Err()
	return err
}

/**
 * Pops item from queue list
 */
func (manager *CacheManager) PopItem(queueKey string) (interface{}, error) {
	return manager.Client.LPop(queueKey).Result()
}

/**
 * Gets
 */
func (manager *CacheManager) ZRevRangeByScore(key string, startScore float64, intPageSize int64) ([]redis.Z, error) {

	var max string
	if startScore != -1 {
		max = fmt.Sprintf("(%.f", startScore)
	} else {
		max = "+inf"
	}
	return manager.Client.ZRevRangeByScoreWithScores(key, redis.ZRangeBy{
		Max:    max,
		Offset: 0,
		Count:  intPageSize,
	}).Result()
}

/**
 * Gets Count of record by member in Sorted List
 */
func (manager *CacheManager) ZGetCountByMember(key string, member string) int64 {
	var score = manager.Client.ZScore(key, member).Val()

	return manager.Client.ZCount(key, util.FloatToString(score), "+inf").Val()
}

/**
 * Gets a item from cache
 */
func (manager *CacheManager) Get(key string) string {
	value, err := manager.Client.Get(key).Result()
	if err != nil && !strings.Contains(err.Error(), "redis: nil") {
		logger.Error(err.Error())
	}
	return value
}

/**
 * Gets a item from cache with error
 */
func (manager *CacheManager) GetWithError(key string) (string, error) {
	value, err := manager.Client.Get(key).Result()
	if err != nil && !strings.Contains(err.Error(), "redis: nil") {
		logger.Error(err.Error())
	}

	return value, err
}

/**
 * Sets a item to cache
 */
func (manager *CacheManager) Set(key string, object interface{}, expireIn time.Duration) {
	out, err := json.Marshal(object)
	if err != nil && !strings.Contains(err.Error(), "redis: nil") {
		logger.Error(err.Error())
	}

	err = manager.Client.Set(key, out, expireIn).Err()
	if err != nil {
		logger.Error(err.Error())
	}
}

// Increments the number stored at key by one
func (manager *CacheManager) Incr(key string) int64 {
	return manager.Client.Incr(key).Val()
}

// Decrements the number stored at key by one
func (manager *CacheManager) Decr(key string) int64 {
	return manager.Client.Decr(key).Val()
}

/**
 *
 */
func (manager *CacheManager) RPush(key, value string) error {
	err := manager.Client.RPush(key, value).Err()
	if err != nil && !strings.Contains(err.Error(), "redis: nil") {
		logger.Error(err.Error())
	}
	return err
}

/**
 *
 */
func (manager *CacheManager) LPop(key string) (string, error) {
	return manager.Client.LPop(key).Result()
}

func (manager *CacheManager) LRange(key string, start, stop int64) []string {
	return manager.Client.LRange(key, start, stop).Val()
}

/**
 *
 */
func (manager *CacheManager) ZCount(key string, min string, max string) int64 {
	res, _ := manager.Client.ZCount(key, min, max).Result()
	return res
}

/**
 *
 */
func (manager *CacheManager) LGetFirst(key string) interface{} {
	return manager.Client.LRange(key, 0, 0)
}

/**
 *
 */
func (manager *CacheManager) LGetAll(key string) []string {
	return manager.Client.LRange(key, 0, -1).Val()
}

/**
 *
 */
func (manager *CacheManager) HSet(key string, field string, value interface{}) error {
	return manager.Client.HSet(key, field, value).Err()
}

/**
 *
 */
func (manager *CacheManager) HGet(key string, field string) (string, error) {
	val, err := manager.Client.HGet(key, field).Result()
	return val, err
}

func (manager *CacheManager) HMGet(key string, fields ...string) (map[string]string, error) {
	val, err := manager.Client.HMGet(key, fields...).Result()
	res := make(map[string]string)
	for i := 0; i < len(fields); i++ {
		if val[i] == nil {
			res[fields[i]] = ""
		} else {
			res[fields[i]] = val[i].(string)
		}
	}
	return res, err
}

/**
 *
 */
func (manager *CacheManager) HMSet(key string, fields map[string]interface{}) error {
	err := manager.Client.HMSet(key, fields).Err()
	return err
}

/**
 *
 */
func (manager *CacheManager) HGetAll(key string) (err error, res map[string]string) {
	res, err = manager.Client.HGetAll(key).Result()
	if err != nil {
		return err, res
	}
	return nil, res
}

/*
* Add item to sorted list
 */
func (client *CacheManager) ZAddItem(key string, score float64, member interface{}) (interface{}, error) {
	return client.Client.ZAdd(key, redis.Z{score, member}).Result()
}

func (client *CacheManager) ZIncrBy(key string, increment float64, member string) float64 {
	return client.Client.ZIncrBy(key, increment, member).Val()
}

/*
* Get elements total of sorted list
 */
func (client *CacheManager) ZGetTotal(key string) (int64, error) {
	return client.Client.ZCard(key).Result()
}

/*
* Get element with highest score
 */
func (client *CacheManager) ZGetHighestScore(key string) ([]redis.Z, error) {
	return client.Client.ZRevRangeWithScores(key, 0, 0).Result()
}

/*
* Get Score of Member in Sorted List
 */
func (client *CacheManager) ZGetScore(key string, member string) float64 {
	return client.Client.ZScore(key, member).Val()
}

func (client *CacheManager) ZGetSortedList(key string, start int64, stop int64) ([]string, error) {
	return client.Client.ZRange(key, start, stop).Val(), nil
}

/**
 * Delete key in sorted list
 */
func (client *CacheManager) ZRem(key string, members string) int64 {
	return client.Client.ZRem(key, members).Val()
}

/**
 * Delete key in sorted list
 */
func (client *CacheManager) HExists(key string, fields string) bool {
	return client.Client.HExists(key, fields).Val()
}

/**
 * Delete key in sorted list
 */
func (client *CacheManager) Exists(key string) int64 {
	return client.Client.Exists(key).Val()
}

/**
 * Delete key in sorted list
 */
func (client *CacheManager) Expire(key string, expiration time.Duration) error {
	return client.Client.Expire(key, expiration).Err()
}

/**
 * Delete key in sorted list
 */
func (client *CacheManager) HDel(key string, field string) error {
	return client.Client.HDel(key, field).Err()
}

/**
 * Delete key in sorted list
 */
func (client *CacheManager) ZRangeByScore(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	vals, err := client.Client.ZRangeByScoreWithScores(key, opt).Result()
	return vals, err
}

func (client *CacheManager) ZRangeLimit1(key string) (string, error) {
	opt := redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  1,
	}
	vals, err := client.Client.ZRangeByScoreWithScores(key, opt).Result()
	if len(vals) == 0 {
		return "", nil
	}
	str := fmt.Sprintf("%v", vals[0].Member)
	return str, err
}

/**
 * Delete key in sorted list
 */
func (client *CacheManager) ZGetSortedSetWithScores(key string, start int64, limit int64) ([]redis.Z, error) {
	var min, max string
	if start == 0 {
		min = util.Int64ToString(start)
		max = "+inf"
	} else {
		max = fmt.Sprintf("(%s", util.Int64ToString(start))
		min = "-inf"
	}
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: 0,
		Count:  limit,
	}
	val, err := client.Client.ZRevRangeByScoreWithScores(key, opt).Result()
	return val, err
}
