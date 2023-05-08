package myredigo

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type (
	RedisClient struct {
		pool *redis.Pool
	}
)

func NewRedisClient(server, password string, db int) *RedisClient {
	return &RedisClient{
		pool: &redis.Pool{
			MaxIdle:     10,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server)
				if err != nil {
					return nil, err
				}
				if password != "" {
					if _, err := c.Do("AUTH", password); err != nil {
						c.Close()
						return nil, err
					}
				}
				if db != 0 {
					if _, err := c.Do("SELECT", db); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

// 定义针对具体业务的方法
func (rc *RedisClient) SetValue(key, value string) error {
	conn := rc.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return fmt.Errorf("failed to set value for key %s: %w", key, err)
	}
	return nil
}

func (rc *RedisClient) GetValue(key string) (string, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", fmt.Errorf("failed to get value for key %s: %w", key, err)
	}
	return value, nil
}

func Init() {
	redisClient := NewRedisClient("127.0.0.1:6379", "123456", 0)
	if err := redisClient.SetValue("name", "jackson"); err != nil {
		log.Fatal(err)
	}
	value, err := redisClient.GetValue("name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	//conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer conn.Close()
	//
	//// 设置键值对
	//_, err = conn.Do("SET", "key", "value")
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 读取键值对
	//value, err := redis.String(conn.Do("GET", "key"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(value)
	//
	//// 关闭 Redis 连接
	//conn.Do("QUIT")

}
