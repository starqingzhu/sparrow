package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func (rc *Redis) SAdd(key string, values ...interface{}) (int64, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	args := redis.Args{}.Add(key).AddFlat(values)

	ret, err := redis.Int64(conn.Do("SADD", args...))
	log.Printf("SAdd key:%s, value:%s, reply:%d, err:%v\n", key, args, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to sadd value for key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) SCard(key string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Int64(conn.Do("SCARD", key))
	log.Printf("SCARD key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to scard for key %s: %w", key, err)
	}

	return ret, err
}

func (r *Redis) SIsMember(key, member string) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Bool(conn.Do("SISMEMBER", key, member))
	log.Printf("SMembers key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to SISMEMBER for key %s: %w", key, err)
	}
	return ret, err
}

func (r *Redis) SMembers(key string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Strings(conn.Do("SMEMBERS", key))
	log.Printf("SMembers key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to SMEMBERS for key %s: %w", key, err)
	}
	return ret, err
}

func (r *Redis) SPop(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.String(conn.Do("SPOP", key))
	log.Printf("SPop key:%s, ret:%s, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to SPop for key %s: %w", key, err)
	}

	return ret, err
}

// 返回集合中的N个随机元素
// 仅仅返回随机元素，而不对集合进行任何改动
func (r *Redis) SRandMember(key string, count int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Strings(conn.Do("SRANDMEMBER", key, count))
	log.Printf("SRandMember key:%s, count:%d, ret:%v, err:%v\n", key, count, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to SISMEMBER for key %s: %w", key, err)
	}
	return ret, err
}

func (r *Redis) SRem(key string, members ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	ret, err := redis.Int64(conn.Do("SREM", members...))
	log.Printf("SRem key:%s, value:%s, reply:%d, err:%v\n", key, members, ret, err)
	if err = checkNil(err); err != nil {
		return ret, fmt.Errorf("failed to SRem value for key %s: %w", key, err)
	}

	return ret, err
}

// // 以下操作注意keys需要hash到相同的slot
//
//	func (r *Redis) SDiff(keys ...string) ([]string, error) {
//		counterIncr("SDiff")
//		startTime := timestrap.Now()
//		success := true
//		defer func() {
//			cmdLatency("SDiff", startTime, success)
//		}()
//		if len(keys) == 0 {
//			success = false
//			return stringsError("SDiff requires at least 1 input keys")
//		}
//		ret, err := r.client.SDiff(r.ctx, keys...).Result()
//		if err = checkNil(err); err != nil {
//			success = false
//		}
//		return ret, err
//	}
//
//	func (r *Redis) SDiffStore(destination string, keys ...string) (int64, error) {
//		counterIncr("SDiffStore")
//		startTime := timestrap.Now()
//		success := true
//		defer func() {
//			cmdLatency("SDiffStore", startTime, success)
//		}()
//		if destination == "" {
//			success = false
//			return int64Error("dest cannot be an empty string")
//		}
//		if len(keys) == 0 {
//			success = false
//			return int64Error("SDiffStore requires at least 1 input keys")
//		}
//		for _, s := range keys {
//			if s == "" {
//				success = false
//				return int64Error("set keys cannot be empty strings")
//			}
//		}
//		ret, err := r.client.SDiffStore(r.ctx, destination, keys...).Result()
//		if err = checkNil(err); err != nil {
//			success = false
//		}
//		return ret, err
//	}
//
//	func (r *Redis) SInter(keys ...string) ([]string, error) {
//		counterIncr("SInter")
//		startTime := timestrap.Now()
//		success := true
//		defer func() {
//			cmdLatency("SInter", startTime, success)
//		}()
//		if len(keys) == 0 {
//			success = false
//			return stringsError("SInter requires at least 1 input keys")
//		}
//		ret, err := r.client.SInter(r.ctx, keys...).Result()
//		if err = checkNil(err); err != nil {
//			success = false
//		}
//		return ret, err
//	}
//
//	func (r *Redis) SInterStore(destination string, keys ...string) (int64, error) {
//		counterIncr("SInterStore")
//		startTime := timestrap.Now()
//		success := true
//		defer func() {
//			cmdLatency("SInterStore", startTime, success)
//		}()
//		if destination == "" {
//			success = false
//			return int64Error("dest cannot be an empty string")
//		}
//		if len(keys) == 0 {
//			success = false
//			return int64Error("SInterStore requires at least 1 input keys")
//		}
//		for _, s := range keys {
//			if s == "" {
//				success = false
//				return int64Error("set keys cannot be empty strings")
//			}
//		}
//		ret, err := r.client.SInterStore(r.ctx, destination, keys...).Result()
//		if err = checkNil(err); err != nil {
//			success = false
//		}
//		return ret, err
//	}
//
//	func (r *Redis) SMove(source, destination string, member interface{}) (bool, error) {
//		counterIncr("SMove")
//		startTime := timestrap.Now()
//		success := true
//		defer func() {
//			cmdLatency("SMove", startTime, success)
//		}()
//		if source == "" || destination == "" {
//			success = false
//			return boolError("key cannot be an empty string")
//		}
//		ret, err := r.client.SMove(r.ctx, source, destination, member).Result()
//		if err = checkNil(err); err != nil {
//			success = false
//		}
//		return ret, err
//	}
//
//	func (r *Redis) SUnion(keys ...string) ([]string, error) {
//		counterIncr("SUnion")
//		startTime := timestrap.Now()
//		success := true
//		defer func() {
//			cmdLatency("SUnion", startTime, success)
//		}()
//		if len(keys) == 0 {
//			success = false
//			return stringsError("SUnion requires at least 1 input set")
//		}
//		for _, s := range keys {
//			if s == "" {
//				success = false
//				return stringsError("set keys cannot be empty strings")
//			}
//		}
//		ret, err := r.client.SUnion(r.ctx, keys...).Result()
//		if err = checkNil(err); err != nil {
//			success = false
//		}
//		return ret, err
//	}
//
//	func (r *Redis) SUnionStore(destination string, keys ...string) (int64, error) {
//		counterIncr("SUnionStore")
//		startTime := timestrap.Now()
//		success := true
//		defer func() {
//			cmdLatency("SUnionStore", startTime, success)
//		}()
//		if destination == "" {
//			success = false
//			return int64Error("dest cannot be an empty string")
//		}
//		if len(keys) == 0 {
//			success = false
//			return int64Error("SUnionStore requires at least 1 input set")
//		}
//		for _, s := range keys {
//			if s == "" {
//				success = false
//				return int64Error("set keys cannot be empty strings")
//			}
//		}
//		ret, err := r.client.SUnionStore(r.ctx, destination, keys...).Result()
//		if err = checkNil(err); err != nil {
//			success = false
//		}
//		return ret, err
//	}
//
