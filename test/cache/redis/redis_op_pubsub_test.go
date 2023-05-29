package redis

func (suite *RedisTestSuite) TestSubscribe() {
	_, err := suite.r.Subscribe("channel3", "channel4")
	suite.Nil(err)

	_, err = suite.r.PSubscribe("channel*", "ch*")
	suite.Nil(err)

	_, err = suite.r.PubSubChannels("channel*")
	suite.Nil(err)

	_, err = suite.r.PUnSubscribe("channel*")
	suite.Nil(err)

	_, err = suite.r.PubSubNumSub("channel1", "channel2")
	suite.Nil(err)

	_, err = suite.r.PubSubNumPat()
	suite.Nil(err)
}
