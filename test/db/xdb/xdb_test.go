package xdb

import (
	"github.com/stretchr/testify/suite"
	"sparrow/pkg/db/xdb"
	"testing"
)

type XDBTestSuite struct {
	suite.Suite
	r *xdb.XDB
}

/*
	User         string `json:"user"`
	Password     string `json:"password"`
	Server       string `json:"server"`
	Port         int32  `json:"port"`
	DataBase     string `json:"database"`
	MaxOpenConns int64  `json:"maxopenconns"`
	MaxIdleConns int64  `json:"maxidleconns"`
*/
// run before the tests in the suite are run.
func (suite *XDBTestSuite) SetupSuite() {
	suite.T().Log("before all tests")
	var conf = xdb.XDbConf{
		User:         "root",
		Password:     "Ab@123456",
		ServerIp:     "192.168.59.177",
		Port:         3306,
		DataBase:     "xdb_user",
		MaxOpenConns: 10,
		MaxIdleConns: 5,
	}
	xdb.Init(&conf)
	suite.r = xdb.Main
}

func (suite *XDBTestSuite) TearDownSuite() {
	suite.T().Log("after all test")
	suite.r.DB.Close()
}

func (suite *XDBTestSuite) SetupTest() {

}

func TestRedisSuite(t *testing.T) {
	suite.Run(t, new(XDBTestSuite))
}
