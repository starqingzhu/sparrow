package redis

//const zKey = "zKey"
//
//func (suite *RedisTestSuite) TestZAdd() {
//	_, err := suite.r.ZAdd(zKey, 1.1, "s1.1", 1.2, "s1.2", 1.3, "s1.3", 1.4, "s1.4", 1.5, "s1.5",
//		1.6, "s1.6", 1.7, "s1.7", 1.8, "s1.8", 1.9, "s1.9", 1.10, "s1.10",
//		1.11, "s1.11", 1.12, "s1.12", 1.13, "s1.13", 1.14, "s1.14")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZAddNX() {
//	var paramList = []interface{}{2.0, "s2.1", 2.1, "s2.2"}
//	_, err := suite.r.ZAddNX(zKey, paramList...)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZAddXX() {
//	var paramList = []interface{}{2.1, "s2.1", 2.2, "s2.2"}
//	_, err := suite.r.ZAddXX(zKey, paramList...)
//	suite.Nil(err)
//}
//func (suite *RedisTestSuite) TestZAddCH() {
//	var paramList = []interface{}{2.1, "s2.1", 2.2, "s2.2"}
//	_, err := suite.r.ZAddCH(zKey, paramList...)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZCard() {
//	_, err := suite.r.ZCard(zKey)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZCount() {
//	_, err := suite.r.ZCount(zKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZIncrBy() {
//	_, err := suite.r.ZIncrBy(zKey, 1, "s2.1")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRange() {
//	_, err := suite.r.ZRange(zKey, 1, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRangeWithScores() {
//	arr, err := suite.r.ZRangeWithScores(zKey, 1, -1)
//	suite.Nil(err)
//	ret := redis.ZUnitArr2Map(arr)
//	log.Println(ret)
//}
//
//func (suite *RedisTestSuite) TestZRangeByScore() {
//	_, err := suite.r.ZRangeByScore(zKey, -math.MaxFloat64, math.MaxFloat64, 0, 100)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRangeByScoreWithScores() {
//	arr, err := suite.r.ZRangeByScoreWithScores(zKey, -math.MaxFloat64, math.MaxFloat64, 0, 100)
//	suite.Nil(err)
//	ret := redis.ZUnitArr2Map(arr)
//	log.Println(ret)
//}
//
//func (suite *RedisTestSuite) TestZRank() {
//	_, err := suite.r.ZRank(zKey, "s2.1")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRem() {
//	_, err := suite.r.ZRem(zKey, "s2.1", "s1.3")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRemRangeByRank() {
//	_, err := suite.r.ZRangeWithScores(zKey, 0, -1)
//	suite.Nil(err)
//	_, err = suite.r.ZRemRangeByRank(zKey, 0, 0)
//	suite.Nil(err)
//	_, err = suite.r.ZRange(zKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRemRangeByScore() {
//	_, err := suite.r.ZRangeWithScores(zKey, 0, -1)
//	suite.Nil(err)
//	_, err = suite.r.ZRemRangeByScore(zKey, 0, 0)
//	suite.Nil(err)
//	_, err = suite.r.ZRange(zKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRevRange() {
//	_, err := suite.r.ZRevRange(zKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRevRangeWithScores() {
//	_, err := suite.r.ZRevRangeWithScores(zKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRevRangeByScore() {
//	_, err := suite.r.ZRevRangeByScore(zKey, math.MaxFloat64, -math.MaxFloat64, 0, 100)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRevRangeByScoreWithScores() {
//	_, err := suite.r.ZRevRangeByScoreWithScores(zKey, math.MaxFloat64, -math.MaxFloat64, 0, 100)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZRevRank() {
//	_, err := suite.r.ZRevRank(zKey, "s2.2")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZScore() {
//	_, err := suite.r.ZScore(zKey, "s2.2")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestZScan() {
//	_, err := suite.r.ZScan(zKey, 0, "s*", 10)
//	suite.Nil(err)
//}
