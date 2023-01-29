package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"github.com/gomodule/redigo/redis"
	"simple-demo/config"
	"time"
)

var (
	redisClient *redis.Pool
	rs          *redsync.Redsync
	//锁过期时间
	lockExpiry = 2 * time.Second
	//获取锁失败重试时间间隔
	retryDelay = 500 * time.Millisecond
	//值过期时间
	valueExpire  = 86400
	ErrMissCache = errors.New("miss Cache")
	//锁设置
	option = []redsync.Option{
		redsync.WithExpiry(lockExpiry),
		redsync.WithRetryDelay(retryDelay),
	}
)

func RedisInit() {
	config := config.Config
	redisConfig := config.RedisConfig
	address := redisConfig.Address
	port := redisConfig.Port
	network := redisConfig.Network
	host := fmt.Sprintf("%s:%s", address, port)
	redisClient = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, host,
				redis.DialDatabase(0),
			)
			if err != nil {
				panic("conn redis failed,")
				return nil, err
			}
			return c, err
		},
	}
	sync := redigo.NewPool(redisClient)
	rs = redsync.New(sync)

}

// ===============string接口=============================
func InsertStringCache(key string, data interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value, "EX", valueExpire)
	if err != nil {
		return err
	}
	return nil
}

func GetStringCache(key string) (interface{}, error) {
	coon := redisClient.Get()
	defer coon.Close()
	value, err := coon.Do("Get", key)
	if err != nil {
		return nil, err
	}
	return value, nil
}
func CloseRedis() {
	redisClient.Close()
}
