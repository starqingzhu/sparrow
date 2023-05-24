package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func (r *Redis) Decr(key string) error {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := conn.Do("DECR", key)
	log.Printf("Decr key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to Decr string for key %s: %w", key, err)
	}
	return nil
}

func (r *Redis) Decrby(key string, increment int64) error {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := conn.Do("DECRBY", key, increment)
	log.Printf("Decrby key:%s, val:%d, ret:%v, err:%v\n", key, increment, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to Decrby string for key %s: %w", key, err)
	}
	return nil
}

func (r *Redis) Get(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if key == "" {
		return "", fmt.Errorf("failed to get value key is nil")
	}
	value, err := redis.String(conn.Do("GET", key))
	if err = checkNil(err); err != nil {
		return "", fmt.Errorf("failed to get value for key %s: %w", key, err)
	}
	return value, nil
}

func (r *Redis) Incr(key string) error {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := conn.Do("INCR", key)
	log.Printf("Incr key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to Incr string for key %s: %w", key, err)
	}
	return nil
}

func (r *Redis) Incrby(key string, increment int64) error {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := conn.Do("INCRBY", key, increment)
	log.Printf("Incrby key:%s, val:%d, ret:%v, err:%v\n", key, increment, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to Incrby string for key %s: %w", key, err)
	}
	return nil
}

func (r *Redis) MGet(keys ...interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()

	if len(keys) == 0 {
		return fmt.Errorf("MGet keys is empty")
	}

	//args := redis.Args{}.AddFlat(keys)
	ret, err := redis.Strings(conn.Do("MGET", keys...))
	log.Printf("MGet keys:%v, ret:%v, err:%v\n", keys, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to MGet args:%v: %w", keys, err)
	}
	return nil
}

func (r *Redis) MSet(values ...interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()

	if len(values) == 0 {
		return fmt.Errorf("MSet values is empty")
	}

	//args := redis.Args{}.AddFlat(values)
	ret, err := conn.Do("MSET", values...)
	log.Printf("MSet args:%v, ret:%v, err:%v\n", values, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to MSet args:%v: %w", values, err)
	}
	return nil
}

func (r *Redis) Set(key, value string) error {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := conn.Do("SET", key, value)
	log.Printf("Set key:%s, val:%s, ret:%v, err:%v\n", key, value, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to set string for key %s: %w", key, err)
	}
	return nil
}

func (r *Redis) SetEX(key string, timeout int64, value string) error {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := conn.Do("SETEX", key, timeout, value)
	log.Printf("SetEX key:%s, timeout:%d, val:%s, ret:%v, err:%v\n", key, timeout, value, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to set SetEX for key %s: %w", key, err)
	}
	return nil
}

func (r *Redis) SetNX(key string, value string) error {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := conn.Do("SETNX", key, value)
	log.Printf("SetNX key:%s, val:%s, ret:%v, err:%v\n", key, value, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to set SetNX for key %s: %w", key, err)
	}
	return nil
}

func (r *Redis) MSetNX(values ...interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()

	if len(values) == 0 {
		return fmt.Errorf("MSetNX values is empty")
	}

	//args := redis.Args{}.AddFlat(values)
	ret, err := conn.Do("MSETNX", values...)
	log.Printf("MSet args:%v, ret:%v, err:%v\n", values, ret, err)
	if err = checkNil(err); err != nil {
		return fmt.Errorf("failed to MSetNX args:%v: %w", values, err)
	}
	return nil
}
