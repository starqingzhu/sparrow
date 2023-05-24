package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

type ZUnit struct {
	Score  float64
	Member interface{}
}

/*
ZADD 是 Redis Sorted Set 数据类型中的一个命令，用于将一个或多个成员加入到 Sorted Set 中。

该命令的语法如下：

ZADD key [NX|XX] [CH] [INCR] score member [score member ...]
其中，

key 是 Sorted Set 的键名。
NX 和 XX 表示控制当成员已经存在时的操作行为，具体解释请见后面的说明。
CH 表示将结果返回给客户端。
INCR 表示将已经存在于 Sorted Set 中的成员的分数进行加减操作。
score 是成员的权重值，可以是整数或浮点数（即支持精度比较高的计算）。
member 是成员的名称，可以是任何字符串。
NX 和 XX 的区别：

NX 表示只有当成员不存在于 Sorted Set 中时才执行操作，如果成员已经存在，则不执行任何操作。
XX 表示只有当成员已经存在于 Sorted Set 中时才执行操作，如果成员不存在，则不执行任何操作。
*/
func (r *Redis) ZAdd(key string, params ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	if _, err := checkParamEven(params...); err != nil {
		return int64Error(err.Error())
	}
	//arr, err := paramsToZArr(params...)
	//if err = checkNil(err); err != nil {
	//	return int64Error(err.Error())
	//}
	args := redis.Args{}.Add(key).AddFlat(params)

	//var ret int64
	ret, err := redis.Int64(conn.Do("ZADD", args...))
	log.Printf("ZAdd key:%s, value:%v, reply:%d, err:%v\n", key, params, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZAdd value for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZAddNX(key string, params ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	if _, err := checkParamEven(params...); err != nil {
		return int64Error(err.Error())
	}

	args := redis.Args{}.Add(key).Add("NX").AddFlat(params)
	ret, err := redis.Int64(conn.Do("ZAdd", args...))
	log.Printf("ZAddNX key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZAddNX for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) ZAddXX(key string, params ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	if _, err := checkParamEven(params...); err != nil {
		return int64Error(err.Error())
	}

	args := redis.Args{}.Add(key).Add("XX").AddFlat(params)
	ret, err := redis.Int64(conn.Do("ZAdd", args...))
	log.Printf("ZAddXX key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZAddXX for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) ZAddCH(key string, params ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	if _, err := checkParamEven(params...); err != nil {
		return int64Error(err.Error())
	}

	args := redis.Args{}.Add(key).Add("CH").AddFlat(params)
	ret, err := redis.Int64(conn.Do("ZAdd", args...))
	log.Printf("ZAddCH key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZAddCH for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) ZCard(key string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("ZCARD", key))
	log.Printf("ZCard key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZCard for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZCount(key string, min int64, max int64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("ZCOUNT", key, min, max))
	log.Printf("ZCount key:%s, ret:%d, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZCount for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZIncrBy(key string, increment float64, member string) (float64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return float64Error(err.Error())
	}

	if err := checkParamString(member); err != nil {
		return float64Error(err.Error())
	}

	ret, err := redis.Float64(conn.Do("ZINCRBY", key, increment, member))
	log.Printf("ZIncrBy key:%s, incrment:%f, member:%s, ret:%f, err:%v\n", key, increment, member, ret, err)
	if err = checkNil(err); err != nil {
		return float64Error(fmt.Sprintf("ZIncrBy to ZIncrBy for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) ZRange(key string, start, stop int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(err.Error())
	}

	ret, err := redis.Strings(conn.Do("ZRANGE", key, start, stop))
	log.Printf("ZRange key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to ZRange for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) ZRangeWithScores(key string, start, stop int64) ([]*ZUnit, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return nil, err
	}

	ret, err := redis.Strings(conn.Do("ZRANGE", key, start, stop, "WITHSCORES"))
	log.Printf("ZRange key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return nil, err
	}

	arr := make([]*ZUnit, 0, len(ret))
	for i, v := range ret {
		if i%2 != 0 {
			continue
		}

		score, parseErr := strconv.ParseFloat(ret[i+1], 10)
		if parseErr != nil {
			return nil, err
		}
		u := &ZUnit{
			Score:  score,
			Member: v,
		}
		arr = append(arr, u)

	}

	return arr, err
}

func (r *Redis) ZRangeByScore(key string, min, max float64, offset, count int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return nil, err
	}

	ret, err := redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count))
	log.Printf("ZRangeByScore key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to ZRangeByScore for key %s: %w", key, err))
	}
	return ret, nil
}

func (r *Redis) ZRangeByScoreWithScores(key string, min, max float64, offset, count int64) ([]*ZUnit, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return nil, err
	}

	ret, err := redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", offset, count))
	log.Printf("ZRangeByScoreWithScores key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return nil, err
	}

	arr := make([]*ZUnit, 0, len(ret))
	for i, v := range ret {
		if i%2 != 0 {
			continue
		}

		score, parseErr := strconv.ParseFloat(ret[i+1], 10)
		if parseErr != nil {
			return nil, err
		}
		u := &ZUnit{
			Score:  score,
			Member: v,
		}
		arr = append(arr, u)

	}

	return arr, err
}

func (r *Redis) ZRank(key, member string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("ZRANK", key, member))
	log.Printf("ZRank key:%s, member:%s, ret:%d, err:%v\n", key, member, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZRank for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZRem(key string, params ...interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}
	if len(params) == 0 {
		return int64Error("invalid param len")
	}

	args := redis.Args{}.Add(key).AddFlat(params)
	ret, err := redis.Int64(conn.Do("ZREM", args...))
	log.Printf("ZRem key:%s, member:%v, ret:%d, err:%v\n", key, params, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZRem for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZRemRangeByRank(key string, start int64, stop int64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("ZREMRANGEBYRANK", key, start, stop))
	log.Printf("ZRemRangeByRank key:%s, start:%d, stop:%d, ret:%d, err:%v\n", key, start, stop, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZRemRangeByRank for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZRemRangeByScore(key string, min float64, max float64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("ZREMRANGEBYSCORE", key, min, max))
	log.Printf("ZRemRangeByScore key:%s, min:%.2f, max:%.2f, ret:%d, err:%v\n", key, min, max, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZRemRangeByScore for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZRevRange(key string, start, stop int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Strings(conn.Do("ZREVRANGE", key, start, stop))
	log.Printf("ZRevRange key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to ZRevRange for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) ZRevRangeWithScores(key string, start, stop int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Strings(conn.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
	log.Printf("ZRevRangeWithScores key:%s, start:%d, stop:%d, ret:%v, err:%v\n",
		key, start, stop, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to ZRevRangeWithScores for key %s: %w", key, err))
	}

	return ret, err

}

func (r *Redis) ZRevRangeByScore(key string, min, max float64, offset, count int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return nil, err
	}

	ret, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, min, max, "LIMIT", offset, count))
	log.Printf("ZRevRangeByScore key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return stringsError(fmt.Sprintf("failed to ZRevRangeByScore for key %s: %w", key, err))
	}
	return ret, nil
}

func (r *Redis) ZRevRangeByScoreWithScores(key string, min, max float64, offset, count int64) ([]*ZUnit, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return nil, err
	}

	ret, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", offset, count))
	log.Printf("ZRevRangeByScoreWithScores key:%s, ret:%v, err:%v\n", key, ret, err)
	if err = checkNil(err); err != nil {
		return nil, err
	}

	arr := make([]*ZUnit, 0, len(ret))
	for i, v := range ret {
		if i%2 != 0 {
			continue
		}

		score, parseErr := strconv.ParseFloat(ret[i+1], 10)
		if parseErr != nil {
			return nil, err
		}
		u := &ZUnit{
			Score:  score,
			Member: v,
		}
		arr = append(arr, u)

	}

	return arr, err
}

func (r *Redis) ZRevRank(key, member string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return int64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Int64(conn.Do("ZREVRANK", key, member))
	log.Printf("ZRevRank key:%s, member:%s, ret:%d, err:%v\n", key, member, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to ZRevRank for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZScore(key, member string) (float64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return float64Error(fmt.Sprintf(err.Error()))
	}

	ret, err := redis.Float64(conn.Do("ZSCORE", key, member))
	log.Printf("ZScore key:%s, member:%s, ret:%.2f, err:%v\n", key, member, ret, err)
	if err = checkNil(err); err != nil {
		return float64Error(fmt.Sprintf("failed to ZScore for key %s: %w", key, err))
	}

	return ret, err
}

func (r *Redis) ZScan(key string, cursor int64, match string, count int64) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if err := checkKey(key); err != nil {
		return stringsError(fmt.Sprintf(err.Error()))
	}

	var ret []string
	vals, err := redis.Values(conn.Do("ZSCAN", key, cursor, "MATCH", match, "COUNT", count))
	if err = checkNil(err); err != nil {
		fmt.Printf("Scan cmd is wrong, err:%s\n", err.Error())
		return ret, fmt.Errorf("failed to Scan keys %w", err)
	}

	cursor, _ = redis.Int64(vals[0], nil)
	ret, _ = redis.Strings(vals[1], nil)
	log.Printf("Scan cursor:%d, match:%s, count:%d, reply:%v, err:%v\n", cursor, match, count, ret, err)
	return ret, err
}

func paramsToZArr(params ...interface{}) ([]interface{}, error) {
	pLen := len(params)
	arr := make([]interface{}, 0, pLen/2)
	for i := range params {
		if i%2 != 0 {
			continue
		}
		var score float64
		var ok bool
		if score, ok = params[i].(float64); !ok {
			var scoreStr string
			if s, ok := params[i].(string); !ok {
				scoreStr = fmt.Sprintf("%v", params[i])
			} else {
				scoreStr = s
			}
			var err error
			if score, err = strconv.ParseFloat(scoreStr, 64); err != nil {
				return nil, err
			}
		}

		z := &ZUnit{
			Score:  score,
			Member: params[i+1],
		}
		arr = append(arr, z)
	}
	return arr, nil
}

func ZUnitArr2Map(param []*ZUnit) map[string]float64 {
	ret := make(map[string]float64, 0)

	for _, v := range param {
		ret[v.Member.(string)] = v.Score
	}
	return ret
}
