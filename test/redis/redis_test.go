package redis

import (
	"github.com/stretchr/testify/suite"
	"sparrow/pkg/cache/redis"
	"testing"
	"time"
)

type RedisTestSuite struct {
	suite.Suite
	r *redis.Redis
}

// run before the tests in the suite are run.
func (suite *RedisTestSuite) SetupSuite() {
	suite.T().Log("before all tests")
	var conf = redis.RedisConfig{
		RedisAddrs:       "192.168.59.177:1000",
		RedisPasswd:      "",
		RedisDBIndex:     0,
		RedisMaxIdle:     3,
		RedisMaxActive:   10,
		RedisIdleTimeout: 240 * time.Second,
	}
	redis.InitMainRedis(conf)

	suite.r = redis.Main
}

func (suite *RedisTestSuite) TearDownSuite() {
	suite.T().Log("after all test")
}

func (suite *RedisTestSuite) SetupTest() {

}

func TestRedisSuite(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}
