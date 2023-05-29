package redis

const lKey = "lKey"

//func (suite *RedisTestSuite) TestLLen() {
//	_, err := suite.r.LLen(lKey)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLPush() {
//	_, err := suite.r.LPush(lKey, "l1", "l2")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLIndex() {
//	_, err := suite.r.LIndex(lKey, 0)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLInsert() {
//	_, err := suite.r.LLen(lKey)
//	suite.Nil(err)
//	_, err = suite.r.LInsertBefore(lKey, "l1", "l0")
//	suite.Nil(err)
//	_, err = suite.r.LLen(lKey)
//	suite.Nil(err)
//
//	_, err = suite.r.LInsertAfter(lKey, "l2", "l3")
//	suite.Nil(err)
//	_, err = suite.r.LLen(lKey)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLPop() {
//	_, err := suite.r.LLen(lKey)
//	suite.Nil(err)
//
//	_, err = suite.r.LPop(lKey)
//	suite.Nil(err)
//
//	_, err = suite.r.LLen(lKey)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLPushX() {
//	_, err := suite.r.LLen(lKey)
//	suite.Nil(err)
//
//	_, err = suite.r.LPushX(lKey, "lx1")
//	suite.Nil(err)
//
//	_, err = suite.r.LLen(lKey)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLRange() {
//	_, err := suite.r.LRange(lKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLRem() {
//	_, err := suite.r.LRange(lKey, 0, -1)
//	suite.Nil(err)
//	_, err = suite.r.LRem(lKey, 10, "l1")
//	suite.Nil(err)
//	_, err = suite.r.LRange(lKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLSet() {
//	_, err := suite.r.LRange(lKey, 0, -1)
//	suite.Nil(err)
//	_, err = suite.r.LSet(lKey, 0, "lset3")
//	suite.Nil(err)
//	_, err = suite.r.LRange(lKey, 0, -1)
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestLTrim() {
//	//_, err := suite.r.LRange(lKey, 0, -1)
//	//suite.Nil(err)
//	//_, err = suite.r.LTrim(lKey, 0, 1)
//	//suite.Nil(err)
//	//_, err = suite.r.LRange(lKey, 0, -1)
//	//suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestBLPop() {
//	var timeOut int64 = 1
//	_, err := suite.r.BLPop(timeOut, lKey, "lKey2")
//	suite.Nil(err)
//
//	_, err = suite.r.BRPop(timeOut, lKey, "lKey2")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestBRPopLPush() {
//	var timeOut int64 = 1
//	_, err := suite.r.BRPopLPush(timeOut, lKey, "lKey2")
//	suite.Nil(err)
//}
//
//func (suite *RedisTestSuite) TestRPopLPush() {
//	_, err := suite.r.RPopLPush(lKey, "lKey2")
//	suite.Nil(err)
//}
