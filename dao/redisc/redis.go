package redisc

import (
	"fmt"
	"gin-test-demo/conf"
	"time"


	"github.com/garyburd/redigo/redis"
	"github.com/go-kratos/kratos/pkg/log"
)

// RedisConnPool .
var RedisConnPool *redis.Pool

// InitRedis .
func InitRedis() {
	cfg := conf.Conf
	maxIdle := cfg.Redis.Idle
	idleTimeout := 240 * time.Second
	_,_ =Test()
	RedisConnPool = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: Test,
		TestOnBorrow: PingRedis,
	}
	fmt.Println("ok" )
}

// PingRedis .
func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		fmt.Print("ping failed")
		log.Error("ping redis fail: %v", err)
	}
	fmt.Print("ping succeed")
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
func Test() (redis.Conn, error) {
	cfg := conf.Conf

	if !cfg.Redis.Enable {
		return nil, nil
	}

	addr := cfg.Redis.Addr
	pass := cfg.Redis.Pass


	connTimeout := time.Duration(cfg.Redis.Timeout.Conn) * time.Millisecond
	readTimeout := time.Duration(cfg.Redis.Timeout.Read) * time.Millisecond
	writeTimeout := time.Duration(cfg.Redis.Timeout.Write) * time.Millisecond
	fmt.Println(connTimeout,readTimeout,writeTimeout)
	fmt.Printf("%v,%v\n", addr, pass)
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("ok, tail:%v\n", err)
		return nil, err
	}

	if pass != "" {
		if _, err := c.Do("AUTH", pass); err != nil {
			c.Close()
			fmt.Printf("ok, tail:%v\n", err)
			return nil, err
		}
	}
	_,err = c.Do("SET","MYKEY","LLFXWZ")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	fmt.Println("ok, tail")
	username, err := redis.String(c.Do("GET", "MYKEY"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	return c, err
}