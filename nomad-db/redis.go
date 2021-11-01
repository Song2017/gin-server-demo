package apiserver

import (
	utils "apiserver/v1/nomad-utils"
	"sync"
	"sync/atomic"

	"github.com/go-redis/redis"
)

var (
	redisMu          sync.Mutex
	redisInitialized uint32
	redisInstance    redis.Client
)

const RedisNil = redis.Nil

func GetRedis() *redis.Client {
	if atomic.LoadUint32(&redisInitialized) == 1 {
		return &redisInstance
	}

	redisMu.Lock() // <-- Lock
	defer redisMu.Unlock()

	if redisInitialized == 0 {
		redisInstance = *redis.NewClient(&redis.Options{
			Addr:     utils.AppConfig().RedisHost,
			Password: utils.AppConfig().RedisPass, // no password set
			PoolSize: 20,

			ReadTimeout: utils.GetTimeOutSeconds(),
		})
		atomic.StoreUint32(&redisInitialized, 1)
	}
	return &redisInstance
}

// func GetRedisPilot() *redis.Client {
// 	once_pilot.Do(func() {
// 		redisInstancePilot = *redis.NewClient(&redis.Options{
// 			Addr:     utils.AppConfig().RedisHostPilot,
// 			Password: utils.AppConfig().RedisPassPilot, // no password set
// 			PoolSize: 20,

// 			ReadTimeout: utils.GetTimeOutSeconds(),
// 		})
// 	})
// 	return &redisInstancePilot
// }
