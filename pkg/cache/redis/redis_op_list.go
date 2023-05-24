package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func (r *Redis) BLPop(timeout int64, keys ...string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKeys(keys...); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	args := redis.Args{}.AddFlat(keys).Add(timeout)
	ret, err := redis.Strings(conn.Do("BLPOP", args...))
	log.Printf("BLPop timeout:%d, keys:%v, ret:%s, err:%v\n", timeout, keys, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to BLPop for key %v: %v", keys, err))
	}

	return ret, err
}

func (r *Redis) BRPop(timeout int64, keys ...string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKeys(keys...); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	args := redis.Args{}.AddFlat(keys).Add(timeout)
	ret, err := redis.Strings(conn.Do("BRPOP", args...))
	log.Printf("BRPop timeout:%d, keys:%v, ret:%s, err:%v\n", timeout, keys, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to BRPop for key %v: %v", keys, err))
	}

	return ret, err
}

func (r *Redis) BRPopLPush(timeout int64, source string, destination string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(source); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	args := redis.Args{}.Add(source).Add(destination).Add(timeout)
	ret, err := redis.String(conn.Do("BRPOPLPUSH", args...))
	log.Printf("BRPopLPush timeout:%d, key:%s, ret:%s, err:%v\n", timeout, source, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to BRPopLPush for key %v: %v", source, err))
	}

	return ret, err
}

func (r *Redis) LIndex(key string, index int64) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.String(conn.Do("LINDEX", key, index))
	log.Printf("LIndex key:%s, index:%d, ret:%s, err:%v\n", key, index, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to LIndex for key %s: %v", key, err))
	}

	return ret, err

}
func (r *Redis) LInsertBefore(key, pivot, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("LINSERT", key, "BEFORE", pivot, value))
	log.Printf("LInsertBefore key:%s, pivot:%s, value:%s, ret:%d, err:%v\n", key, pivot, value, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to LInsertBefore for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) LInsertAfter(key, pivot, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("LINSERT", key, "AFTER", pivot, value))
	log.Printf("LInsertAfter key:%s, pivot:%s, value:%s, ret:%d, err:%v\n", key, pivot, value, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to LInsertAfter for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) LLen(key string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("LLEN", key))
	log.Printf("LLen key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to LLen for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) LPop(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.String(conn.Do("LPOP", key))
	log.Printf("LPop key:%s, ret:%s, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to LPop for key %s: %v", key, err))
	}
	return ret, err

}

func (r *Redis) LPush(key string, vs ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	if len(vs) == 0 {
		return int64Error(fmt.Sprintf("params len is 0"))
	}

	args := redis.Args{}.Add(key).AddFlat(vs)
	ret, err := redis.Int64(conn.Do("LPUSH", args...))
	log.Printf("LPush key:%s, params:%v, ret:%d, err:%v\n", key, vs, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to LPush for key %s: %v", key, err))
	}

	return ret, err

}

func (r *Redis) LPushX(key string, value interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("LPUSHX", key, value))
	log.Printf("LPushX key:%s, value:%v, ret:%d, err:%v\n", key, value, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to LPushX for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) LRange(key string, start, stop int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Strings(conn.Do("LRANGE", key, start, stop))
	log.Printf("LRange key:%s, start:%d, stop:%d, ret:%v, err:%v\n", key, start, stop, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to LRange for key %s: %v", key, err))
	}

	return ret, err

}

func (r *Redis) LRem(key string, count int64, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("LREM", key, count, value))
	log.Printf("LRem key:%s, count:%d, value:%s, ret:%d, err:%v\n", key, count, value, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to LRem for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) LSet(key string, index int64, value string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.String(conn.Do("LSET", key, index, value))
	log.Printf("LSet key:%s, index:%d, value:%s, ret:%s, err:%v\n", key, index, value, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to LSet for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) LTrim(key string, start, stop int64) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}
	ret, err := redis.String(conn.Do("LTRIM", key, start, stop))
	log.Printf("LTrim key:%s, start:%d, stop:%d, ret:%s, err:%v\n", key, start, stop, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to LTrim for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) RPop(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.String(conn.Do("RPOP", key))
	log.Printf("RPop key:%s, ret:%s, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to RPop for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) RPopLPush(source string, destination string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(source); err != nil {
		return stringError(fmt.Sprintf(err.Error()))
	}

	args := redis.Args{}.Add(source).Add(destination)
	ret, err := redis.String(conn.Do("RPopLPush", args...))
	log.Printf("RPopLPush source:%s, destination:%s, ret:%s, err:%v\n", source, destination, ret, err)
	if err = checkNil(err); err != nil {
		return stringError(fmt.Sprintf("failed to RPopLPush for key %v: %v", source, err))
	}
	err = nil
	return ret, err
}

func (r *Redis) RPush(key string, vs ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	if len(vs) == 0 {
		return int64Error(fmt.Sprintf("params len is 0"))
	}

	args := redis.Args{}.Add(key).AddFlat(vs)
	ret, err := redis.Int64(conn.Do("RPUSH", args...))
	log.Printf("RPush key:%s, params:%v, ret:%d, err:%v\n", key, vs, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to RPush for key %s: %v", key, err))
	}

	return ret, err
}

func (r *Redis) RPushX(key string, value interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("RPUSHX", key, value))
	log.Printf("RPushX key:%s, value:%v, ret:%d, err:%v\n", key, value, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to RPushX for key %s: %v", key, err))
	}

	return ret, err
}
