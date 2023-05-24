package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func (r *Redis) HSet(key, field, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("HSET", key, field, value))
	log.Printf("HSet key:%s, field:%s, value:%s, ret:%d, err:%v\n", key, field, value, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to HSet for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HSetNX(key, field, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("HSETNX", key, field, value))
	log.Printf("HSetNX key:%s, field:%s, value:%s, ret:%d, err:%v\n", key, field, value, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to HSetNX for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HMSet(key string, params ...interface{}) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	if _, err := checkParamEven(params...); err != nil {
		return stringError(err.Error())
	}

	args := redis.Args{}.Add(key).AddFlat(params)
	ret, err := redis.String(conn.Do("HMSET", args...))
	log.Printf("HMSet key:%s, key-val:%v, ret:%s, err:%v\n", key, params, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to HMSet for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HGet(key string, field string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.String(conn.Do("HGET", key, field))
	log.Printf("HGet key:%s, filed:%s, ret:%s, err:%v\n", key, field, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to HGet for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HMGet(key string, params ...string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	args := redis.Args{}.Add(key).AddFlat(params)
	ret, err := redis.Strings(conn.Do("HMGET", args...))
	log.Printf("HMGet key:%s, keys:%v, ret:%s, err:%v\n", key, params, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to HMGet for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HGetAll(key string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}
	ret, err := redis.Strings(conn.Do("HGETALL", key))
	log.Printf("HGetAll key:%s, ret:%s, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to HMGet for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HIncrBy(key, field string, incr int64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}
	ret, err := redis.Int64(conn.Do("HINCRBY", key, field, incr))
	log.Printf("HIncrBy key:%s, field:%s, incr:%d, ret:%d, err:%v\n", key, field, incr, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to HIncrBy for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HIncrByFloat(key, field string, incr float64) (float64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return float64Error(fmt.Sprintf(err.Error()))
	}
	ret, err := redis.Float64(conn.Do("HINCRBYFLOAT", key, field, incr))
	log.Printf("HIncrByFloat key:%s, field:%s, incr:%.2f, ret:%.2f, err:%v\n", key, field, incr, ret, err)
	if err = checkNil(err); err != nil {
		return float64Error(fmt.Sprintf("failed to HIncrByFloat for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HDel(key string, fields ...string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	args := redis.Args{}.Add(key).AddFlat(fields)
	ret, err := redis.Int64(conn.Do("HDEL", args...))
	log.Printf("HDel key:%s, field:%v, ret:%d, err:%v\n", key, fields, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to HDel for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HExists(key, field string) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return boolError(fmt.Sprintf(err.Error()))
	}
	ret, err := redis.Bool(conn.Do("HEXISTS", key, field))
	log.Printf("HExists key:%s, field:%v, ret:%t, err:%v\n", key, field, ret, err)
	if err = checkNil(err); err != nil {
		return boolError(fmt.Sprintf("failed to HExists for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HKeys(key string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Strings(conn.Do("HKEYS", key))
	log.Printf("HKeys key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to HKeys for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HVals(key string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Strings(conn.Do("HVALS", key))
	log.Printf("HVals key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to HVals for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HLen(key string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("HLEN", key))
	log.Printf("HLen key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to HLen for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) HScan(key string, cursor int64, match string, count int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	var ret []string
	vals, err := redis.Values(conn.Do("HSCAN", key, cursor, "MATCH", match, "COUNT", count))
	if err = checkNil(err); err != nil {
		log.Printf("HScan cmd is wrong, err:%s\n", err.Error())
		return ret, fmt.Errorf("failed to HScan keys %s", err.Error())
	}

	cursor, _ = redis.Int64(vals[0], nil)
	ret, _ = redis.Strings(vals[1], nil)
	log.Printf("HScan cursor:%d, match:%s, count:%d, reply:%v, err:%v\n", cursor, match, count, ret, err)
	return ret, err
}
