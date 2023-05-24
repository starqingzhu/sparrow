package redis

const tempKey = "tempKey"

//func (suite *RedisTestSuite) TestDel() {
//	err := suite.r.Set(tempKey, "tempVal")
//	suite.Nil(err)
//
//	del, err1 := suite.r.Del(tempKey)
//	suite.Equal(del, int64(1))
//	suite.Nil(err1)
//}
//
//func (suite *RedisTestSuite) TestExists() {
//	err := suite.r.Set(tempKey, "tempVal")
//	suite.Nil(err)
//	exist, err := suite.r.Exists(tempKey)
//	suite.Nil(err)
//	suite.True(exist)
//}
