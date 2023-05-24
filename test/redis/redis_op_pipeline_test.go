package redis

const pipeKey = "pipeKey"
const pipeKeyExpire = 60 * 20

//func (suite *RedisTestSuite) TestPipe() {
//
//	cmdList := []interface{}{
//		[]interface{}{"INCRBY", pipeKey, 1},
//		[]interface{}{"EXPIRE", pipeKey, pipeKeyExpire},
//		[]interface{}{"INCRBY", pipeKey, 2},
//		[]interface{}{"INCRBY", pipeKey, 3},
//	}
//
//	err := suite.r.Pipeline(cmdList)
//	suite.Nil(err)
//
//	err = suite.r.Pipeline(cmdList)
//	suite.Nil(err)
//
//	err = suite.r.Pipeline(cmdList)
//	suite.Nil(err)
//
//	err = suite.r.Pipeline(cmdList)
//	suite.Nil(err)
//
//	err = suite.r.Pipeline(cmdList)
//	suite.Nil(err)
//}

//func TestPipe(t *testing.T) {
//	var conf = redis.RedisConfig{
//		RedisAddrs:       "192.168.59.177:1000",
//		RedisPasswd:      "",
//		RedisDBIndex:     0,
//		RedisMaxIdle:     3,
//		RedisMaxActive:   10,
//		RedisIdleTimeout: 240 * time.Second,
//	}
//	redis.InitMainRedis(conf)
//
//	cmdList := []interface{}{
//		[]interface{}{"INCRBY", pipeKey, 1},
//		[]interface{}{"EXPIRE", pipeKey, pipeKeyExpire},
//		[]interface{}{"INCRBY", pipeKey, 2},
//		[]interface{}{"INCRBY", pipeKey, 3},
//	}
//	redis.Main.Pipeline(cmdList)
//}
