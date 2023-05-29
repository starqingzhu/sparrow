package redis

const hashKey = "hashKey"
const hashKeyExpire = 60 * 20

//func (suite *RedisTestSuite) TestHSet() {
//	_, err := suite.r.HSet(hashKey, "field1", "0")
//	suite.Nil(err)
//	_, err = suite.r.HSetNX(hashKey, "field1", "0")
//	suite.Nil(err)
//	_, err = suite.r.HSetNX(hashKey, "field2", "0")
//	suite.Nil(err)
//	suite.r.Expire(hashKey, hashKeyExpire)
//
//	paramsList := []interface{}{"field1", "0", "field2", "0", "field3", "0"}
//	_, err = suite.r.HMSet(hashKey, paramsList...)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestHGet() {
//	_, err := suite.r.HGet(hashKey, "field1")
//	suite.Nil(err)
//
//	paramsList := []string{"field1", "field2"}
//	_, err = suite.r.HMGet(hashKey, paramsList...)
//	suite.Nil(err)
//
//	suite.r.HIncrBy(hashKey, "field1", 2)
//	suite.r.HIncrByFloat(hashKey, "field1", 2.5)
//	suite.r.HDel(hashKey, "field2")
//	suite.r.HExists(hashKey, "field1")
//	suite.r.HKeys(hashKey)
//	suite.r.HVals(hashKey)
//	suite.r.HLen(hashKey)
//	suite.r.HScan(hashKey, 0, "f*", 10)
//
//	_, err = suite.r.HGetAll(hashKey)
//	suite.Nil(err)
//}
