package redisc

import (
	"fmt"
	"time"

	"kops/app/monitor-api/conf"

	"github.com/garyburd/redigo/redis"
	"github.com/go-kratos/kratos/pkg/log"
)

// RedisConnPool .
var RedisConnPool *redis.Pool

// InitRedis .
func InitRedis() {
	cfg := conf.Conf

	if !cfg.Redis.Enable {
		return
	}

	addr := cfg.Redis.Addr
	pass := cfg.Redis.Pass
	maxIdle := cfg.Redis.Idle
	idleTimeout := 240 * time.Second

	connTimeout := time.Duration(cfg.Redis.Timeout.Conn) * time.Millisecond
	readTimeout := time.Duration(cfg.Redis.Timeout.Read) * time.Millisecond
	writeTimeout := time.Duration(cfg.Redis.Timeout.Write) * time.Millisecond

	RedisConnPool = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialConnectTimeout(connTimeout), redis.DialReadTimeout(readTimeout), redis.DialWriteTimeout(writeTimeout))
			if err != nil {
				return nil, err
			}

			if pass != "" {
				if _, err := c.Do("AUTH", pass); err != nil {
					c.Close()
					log.Error("redis auth fail")
					return nil, err
				}
			}
   			fmt.Println("ok" ,err)
			return c, err
		},
		TestOnBorrow: PingRedis,
	}
}

// PingRedis .
func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		log.Error("ping redis fail: %v", err)
	}
	return err
}

// CloseRedis .
func CloseRedis() {
	if !conf.Conf.Redis.Enable {
		return
	}
	log.Info("closing redis...")
	RedisConnPool.Close()
}
