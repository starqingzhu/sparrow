package redis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type (
	RedisConfig struct {
		RedisAddrs       string        `default:"localhost:6379" cfg:"redis_addrs" json:"redis_addrs"`
		RedisPasswd      string        `cfg:"redis_passwd" default:"" json:"redis_passwd"`
		RedisDBIndex     int           `cfg:"redis_db_index" json:"redis_db_index"`
		RedisMaxIdle     int           `default:"128" cfg:"redis_max_idle" json:"redis_max_idle"`
		RedisMaxActive   int           `default:"2048" cfg:"redis_max_active" json:"redis_max_active"`
		RedisIdleTimeout time.Duration `default:"300000" cfg:"redis_idle_timeout" json:"redis_idle_timeout"`
		RedisPoolSize    int           `default:"128" cfg:"redis_pool_size" json:"redis_pool_size"`
	}

	Redis struct {
		conf      RedisConfig
		pool      *redis.Pool
		ctx       context.Context // redis op context
		isCluster bool            // redis cluster mode
	}
)

var Main *Redis

// var redisMap map[string]*Redis
func InitMainRedis(cfg RedisConfig) {
	Main = NewRedis(cfg)
	if Main == nil {
		log.Fatalf("redis.InitMainRedis NewRedis error, cfg:+v\n", cfg)
	}
	log.Printf("redis.InitMainRedis init success\n")
}

func NewRedis(cfg RedisConfig) *Redis {
	return &Redis{
		pool: &redis.Pool{
			MaxIdle:     cfg.RedisMaxIdle,
			MaxActive:   cfg.RedisMaxActive,
			IdleTimeout: cfg.RedisIdleTimeout,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("link", cfg.RedisAddrs)
				if err = checkNil(err); err != nil {
					return nil, err
				}
				if cfg.RedisPasswd != "" {
					if _, err := c.Do("AUTH", cfg.RedisPasswd); err != nil {
						c.Close()
						return nil, err
					}
				}
				if cfg.RedisDBIndex != 0 {
					if _, err := c.Do("SELECT", cfg.RedisDBIndex); err != nil {
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

func (r *Redis) String() string {
	return fmt.Sprintf("%v", r.pool.Stats())
}

func (r *Redis) Close() error {
	return r.pool.Close()
}

// 定义针对具体业务的方法
//func Init() {
//	var conf = RedisConfig{
//		RedisAddrs:       "192.168.59.177:1000",
//		RedisPasswd:      "",
//		RedisDBIndex:     0,
//		RedisMaxIdle:     3,
//		RedisMaxActive:   10,
//		RedisIdleTimeout: 240 * timestrap.Second,
//	}
//	redisClient := NewRedis(conf)
//
//	if redisClient == nil {
//		panic("redis connect nil")
//	}
//
//	fmt.Println("redisClient:", redisClient.String())
//
//	//string
//	var key1 = "name"
//	redisClient.Set(key1, "jackson")
//
//	var key2 = "name2"
//	redisClient.SetEX(key2, 5*60, "jackson2")
//	redisClient.SetNX(key2, "jackson2b")
//
//	var key3 = "name3"
//	redisClient.SetNX(key3, "jackson3")
//
//	value, err := redisClient.Get(key1)
//	fmt.Println(value)
//
//	//set
//	var sKey = "sname"
//	//var retInt64 int64
//	setArr := []interface{}{"s1", "s2", "s3", "s4"}
//	//args := redis.Args{}.AddFlat(setArr)
//	_, err = redisClient.SAdd(
//		sKey, setArr...,
//	)
//	if err = checkNil(err); err != nil {
//		log.Fatal(err)
//	}
//
//	var mKVs = map[string]string{
//		"kv_key1": "kv_val1",
//		"kv_key2": "kv_val2",
//		"kv_key3": "kv_val3",
//		"kv_key4": "kv_val4",
//		"kv_key5": "kv_val5",
//		"kv_key6": "kv_val6",
//	}
//
//	kvArgsArr := make([]interface{}, 0)
//	kvArgsArr2 := make([]interface{}, 0)
//	kArgsArr := make([]interface{}, 0)
//	for k, v := range mKVs {
//		kvArgsArr = append(kvArgsArr, k)
//		kvArgsArr = append(kvArgsArr, v)
//
//		kvArgsArr2 = append(kvArgsArr2, k)
//		kvArgsArr2 = append(kvArgsArr2, v+"xx")
//
//		kArgsArr = append(kArgsArr, k)
//	}
//
//	redisClient.MSet(kvArgsArr...)
//	redisClient.MSetNX(kvArgsArr2...)
//	redisClient.MGet(kArgsArr...)
//
//	countKey1 := "countKey1"
//	redisClient.Incr(countKey1)
//	redisClient.Incrby(countKey1, 2)
//	redisClient.Decr(countKey1)
//	redisClient.Decrby(countKey1, 2)
//
//	//redisClient.SCard(sKey)
//	//redisClient.SPop(sKey)
//	//redisClient.SCard(sKey)
//	//redisClient.SMembers(sKey)
//	//redisClient.SIsMember(sKey, "s1")
//	//redisClient.SRem(sKey, "s2")
//	//redisClient.SRandMember(sKey, 2)
//	//zset
//	zsetKey := "zsetKey"
//	redisClient.ZAdd(zsetKey, float64(1.1), "mem1")
//	//list
//	//hash
//	//key操作
//	//redisClient.Del(key1)
//	redisClient.Exists(sKey)
//	redisClient.ExpireAt(sKey, timestrap.Now().Add(timestrap.Second*(5*60)).Unix())
//	redisClient.Expire(sKey, 5*60)
//	redisClient.TTL(sKey)
//	redisClient.Persist(sKey)
//	redisClient.PTTL(sKey)
//	redisClient.Scan(0, "*", 1)
//
//	redisClient.Close()
//}
