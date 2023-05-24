package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

/*
Redis 发布订阅机制是一种非常实用的消息传递模式，它可以广泛用于各种场景，包括但不限于以下几个方面：
	1)实时通知和消息推送：Redis 发布订阅机制可以用于实时通知和消息推送，例如社交网络、实时聊天、在线游戏和在线竞技场等场景。通过订阅相应的通道，用户可以实时收到新的消息和事件，从而获取最新的动态和信息。
	2)分布式系统协调和通信：Redis 发布订阅机制可以用于分布式系统中的协调和通信，例如分布式锁、任务调度、集群间的通信和状态同步等场景。通过订阅相应的通道，各个节点可以实时获取集群中的状态变化和事件通知，并相应地进行处理。
	3)数据缓存和预热：Redis 发布订阅机制可以用于数据缓存和预热，例如缓存预热和数据更新通知等场景。通过订阅相应的通道，缓存节点可以实时获取数据更新通知和缓存预热请求，并相应地进行处理。
	4)日志收集和统计分析：Redis 发布订阅机制可以用于日志收集和统计分析，例如事件日志和行为数据等场景。通过订阅相应的通道，统计节点可以实时获取事件日志和行为数据，并对其进行统计分析和处理。

总的来说，Redis 发布订阅机制具有很强的扩展性和灵活性，可以适用于各种场景和需求。同时，它还提供了多种高级特性，例如模式匹配、订阅分布和消息持久化等，可以进一步增强其潜力和可靠性。
*/

func (r *Redis) Publish(channel string, message interface{}) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if channel == "" || message == "" {
		return int64Error("channel or message cannot be empty string")
	}

	ret, err := redis.Int64(conn.Do("PUBLISH", channel, message))
	log.Printf("Publish channel:%s, message:%v, ret:%d, err:%v\n", channel, message, ret, err)
	if err = checkNil(err); err != nil {
		return int64Error(fmt.Sprintf("failed to Publish for channel %s: %v", channel, err))
	}
	return ret, err
}

func (r *Redis) Subscribe(channels ...string) (*redis.Subscription, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if len(channels) == 0 {
		return nil, errors.New(fmt.Sprintf("Subscribe channes is empty"))
	}

	args := redis.Args{}.AddFlat(channels)
	ret, err := conn.Do("SUBSCRIBE", args...)
	aRet, ok := ret.([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("Subscribe ret type wrong"))
	}
	pRet := &redis.Subscription{}
	pRet.Kind, _ = redis.String(aRet[0], err)
	pRet.Channel, _ = redis.String(aRet[1], err)
	pRet.Count, _ = redis.Int(aRet[2], err)

	log.Printf("Subscribe channels:%s, ret:%v, pRet:%v, err:%v\n", channels, ret, pRet, err)
	if err = checkNil(err); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to Subscribe for channels %s: %v", channels, err))
	}
	return pRet, err
}

func (r *Redis) PSubscribe(channels ...string) (*redis.Subscription, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if len(channels) == 0 {
		return nil, errors.New(fmt.Sprintf("PSubscribe channes is empty"))
	}

	args := redis.Args{}.AddFlat(channels)
	ret, err := conn.Do("PSUBSCRIBE", args...)
	aRet, ok := ret.([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("PSubscribe ret type wrong"))
	}
	pRet := &redis.Subscription{}
	pRet.Kind, _ = redis.String(aRet[0], err)
	pRet.Channel, _ = redis.String(aRet[1], err)
	pRet.Count, _ = redis.Int(aRet[2], err)

	log.Printf("PSubscribe channels:%s, ret:%v, pRet:%v, err:%v\n", channels, ret, pRet, err)
	if err = checkNil(err); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to PSubscribe for channels %s: %v", channels, err))
	}
	return pRet, err
}

func (r *Redis) UnSubscribe(channels ...string) (*redis.Subscription, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if len(channels) == 0 {
		return nil, errors.New(fmt.Sprintf("UnSubscribe channes is empty"))
	}

	args := redis.Args{}.AddFlat(channels)
	ret, err := conn.Do("UNSUBSCRIBE", args...)
	aRet, ok := ret.([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("UnSubscribe ret type wrong"))
	}
	pRet := &redis.Subscription{}
	pRet.Kind, _ = redis.String(aRet[0], err)
	pRet.Channel, _ = redis.String(aRet[1], err)
	pRet.Count, _ = redis.Int(aRet[2], err)

	log.Printf("UnSubscribe channels:%s, ret:%v, pRet:%v, err:%v\n", channels, ret, pRet, err)
	if err = checkNil(err); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to UnSubscribe for channels %s: %v", channels, err))
	}
	return pRet, err
}

func (r *Redis) PUnSubscribe(channels ...string) (*redis.Subscription, error) {
	conn := r.pool.Get()
	defer conn.Close()

	if len(channels) == 0 {
		return nil, errors.New(fmt.Sprintf("PUnSubscribe channes is empty"))
	}

	args := redis.Args{}.AddFlat(channels)
	ret, err := conn.Do("PUNSUBSCRIBE", args...)
	aRet, ok := ret.([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("PUnSubscribe ret type wrong"))
	}
	pRet := &redis.Subscription{}
	pRet.Kind, _ = redis.String(aRet[0], err)
	pRet.Channel, _ = redis.String(aRet[1], err)
	pRet.Count, _ = redis.Int(aRet[2], err)

	log.Printf("PUnSubscribe channels:%s, ret:%v, pRet:%v, err:%v\n", channels, ret, pRet, err)
	if err = checkNil(err); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to PUnSubscribe for channels %s: %v", channels, err))
	}
	return pRet, err
}

/*
PUBSUB CHANNELS [pattern]
列出当前的活跃频道。

活跃频道指的是那些至少有一个订阅者的频道， 订阅模式的客户端不计算在内。

pattern 参数是可选的：

如果不给出 pattern 参数，那么列出订阅与发布系统中的所有活跃频道。
如果给出 pattern 参数，那么只列出和给定模式 pattern 相匹配的那些活跃频道。
复杂度： O(N) ， N 为活跃频道的数量（对于长度较短的频道和模式来说，将进行模式匹配的复杂度视为常数）。

返回值： 一个由活跃频道组成的列表。
*/
func (r *Redis) PubSubChannels(pattern string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	args := redis.Args{}.Add("CHANNELS")
	if len(pattern) != 0 {
		args.Add(pattern)
	}
	ret, err := redis.Strings(conn.Do("PUBSUB", args...))
	log.Printf("PubSubChannels pattern:%s, ret:%v,  err:%v\n", pattern, ret, err)
	if err = checkNil(err); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to PubSubChannels for pattern %s: %v", pattern, err))
	}
	return ret, err
}

type NumSubRetInfo struct {
	Channel string
	Count   int64
}

/*
返回给定频道的订阅者数量， 订阅模式的客户端不计算在内。

复杂度： O(N) ， N 为给定频道的数量。

返回值： 一个多条批量回复（Multi-bulk reply），回复中包含给定的频道，以及频道的订阅者数量。 格式为：频道 channel-1 ， channel-1 的订阅者数量，频道 channel-2 ， channel-2 的订阅者数量，诸如此类。 回复中频道的排列顺序和执行命令时给定频道的排列顺序一致。 不给定任何频道而直接调用这个命令也是可以的， 在这种情况下， 命令只返回一个空列表。
*/
func (r *Redis) PubSubNumSub(channels ...string) (retArrInfo []NumSubRetInfo, err error) {
	conn := r.pool.Get()
	defer conn.Close()
	if len(channels) == 0 {
		return nil, errors.New(fmt.Sprintf("channels is empty"))
	}
	args := redis.Args{}.Add("NUMSUB").AddFlat(channels)

	var ret interface{}
	ret, err = conn.Do("PUBSUB", args...)
	log.Printf("PubSubNumSub channels:%v, ret:%v,  err:%v\n", channels, ret, err)
	if err = checkNil(err); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to PubSubNumSub for pattern %v: %s", channels, err.Error()))
	}

	aRet, ok := ret.([]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("PubSubNumSub ret type wrong"))
	}

	nLen := len(aRet)
	for i := 0; i < nLen; i += 2 {
		retInfoNode := NumSubRetInfo{}
		retInfoNode.Channel, _ = redis.String(aRet[i], err)
		retInfoNode.Count, _ = redis.Int64(aRet[i+1], err)
		retArrInfo = append(retArrInfo, retInfoNode)
	}
	log.Printf("PubSubNumSub channels:%v, ret:%v, retArrInfo：%v, err:%v\n", channels, ret, retArrInfo, err)

	return retArrInfo, err
}

func (r *Redis) PubSubNumPat() (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()

	args := redis.Args{}.Add("NUMPAT")

	ret, err := redis.Int64(conn.Do("PUBSUB", args...))
	log.Printf("PubSubNumPat , ret:%v,  err:%v\n", ret, err)
	if err = checkNil(err); err != nil {
		return redis.Int64(0, err)
	}

	return ret, err
}
