package redis

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

func (r *RedisClient) String() string {
	return fmt.Sprintf("%v", r.pool.Stats())
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

func (rc *RedisClient) SAdd(key, value string) (int64, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	var ret int64
	reply, err := conn.Do("SADD", key, value)
	log.Printf("SAdd key:%s, value:%s, reply:%d, err:%v\n", key, value, reply, err)
	if err != nil {
		return ret, fmt.Errorf("failed to set value for key %s: %w", key, err)
	}

	ret = reply.(int64)

	return ret, err
}

func (rc *RedisClient) SCard(key string) (int64, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	var ret int64
	reply, err := conn.Do("SCARD", key)
	log.Printf("SCARD key:%s, reply:%d, err:%v\n", key, reply, err)
	if err != nil {
		return ret, fmt.Errorf("failed to set value for key %s: %w", key, err)
	}

	ret = reply.(int64)

	return ret, err
}

func (rc *RedisClient) SPop(key, val string) (int64, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	var ret int64
	reply, err := conn.Do("SCARD", key)
	log.Printf("SPop key:%s, reply:%d, err:%v\n", key, reply, err)
	if err != nil {
		return ret, fmt.Errorf("failed to set value for key %s: %w", key, err)
	}

	ret = reply.(int64)

	return ret, err
}

func Init() {
	redisClient := NewRedisClient("127.0.0.1:6379", "123456", 0)

	if redisClient == nil {
		panic("redis connect nil")
	}

	fmt.Println("redisClient:", redisClient.String())

	//string
	if err := redisClient.SetValue("name", "jackson"); err != nil {
		log.Fatal(err)
	}
	value, err := redisClient.GetValue("name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	//set
	var sKey = "sname"
	//var retInt64 int64
	var setArr = []string{"s1", "s2", "s3", "s4"}
	for _, m := range setArr {
		_, err = redisClient.SAdd(sKey, m)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = redisClient.SCard(sKey)
	//zset
	//list
	//hash

}
