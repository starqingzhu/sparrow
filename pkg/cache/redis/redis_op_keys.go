package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func (r *Redis) Del(keys ...string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	args := redis.Args{}.AddFlat(keys)
	ret, err := redis.Int64(conn.Do("DEL", args...))
	log.Printf("Del keys:%v, reply:%d, err:%v\n", args, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to Del value for keys %s: %w", args, err)
	}

	return ret, err
}

func (r *Redis) Exists(key string) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Bool(conn.Do("EXISTS", key))
	log.Printf("Exists keys:%v, reply:%t, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to Exists key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) ExpireAt(key string, t int64) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Bool(conn.Do("EXPIREAT", key, t))
	log.Printf("ExpireAt key:%s, timeout:%d, reply:%t, err:%v\n", key, t, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to ExpireAt key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) Expire(key string, t int64) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Bool(conn.Do("EXPIRE", key, t))
	log.Printf("EXPIRE key:%s, timeout:%d, reply:%t, err:%v\n", key, t, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to EXPIRE key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) TTL(key string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Int64(conn.Do("TTL", key))
	log.Printf("TTL key:%s, reply:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to TTL key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) PTTL(key string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Int64(conn.Do("PTTL", key))
	log.Printf("PTTL key:%s, reply:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to PTTL key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) Persist(key string) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Bool(conn.Do("PERSIST", key))
	log.Printf("Persist key:%s, reply:%t, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to Persist key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) Scan(cursor int64, match string, count int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	var ret []string
	//if func(c string) error {
	//	var cmdMap = map[string]bool{
	//		"SCAN":   true,
	//		"SSCAN":  true,
	//		"HSCAN":  true,
	//		"ZHSCAN": true,
	//	}
	//	if _, ok := cmdMap[cmd]; !ok {
	//		return fmt.Errorf("cmd is wrong")
	//	}
	//	return nil
	//}(cmd) != nil {
	//	fmt.Printf("Scan cmd is wrong, cmd:%s\n", cmd)
	//	return ret, fmt.Errorf("Scan cmd is wrong, cmd:%s", cmd)
	//}

	vals, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", match, "COUNT", count))
	if err = checkNil(err); err != nil {
		fmt.Printf("Scan cmd is wrong, err:%s\n", err.Error())
		return ret, fmt.Errorf("failed to Scan keys %w", err)
	}

	cursor, _ = redis.Int64(vals[0], nil)
	ret, _ = redis.Strings(vals[1], nil)
	log.Printf("Scan cursor:%d, match:%s, count:%d, reply:%v, err:%v\n", cursor, match, count, ret, err)
	return ret, err
}
