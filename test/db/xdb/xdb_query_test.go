package xdb

import (
	"sparrow/internal/table/tbstruct"
)

func (suite *XDBTestSuite) TestXdbUserQuery() {

	querySql := "select * from user"
	var users []*tbstruct.User

	// 方案1 xdb查询
	rows, err := suite.r.Queryx(querySql)
	suite.Nil(err)
	for rows.Next() {
		user := &tbstruct.User{}
		rows.StructScan(user)
		users = append(users, user)
	}

	// 方案2 官方接口查询
	//rows, err := suite.r.Query(querySql)
	//suite.Nil(err)
	//
	//for rows.Next() {
	//	var user tbstruct.User
	//	arr, err1 := tbstruct.GetObjMemAddrArr(&user)
	//	suite.Nil(err1)
	//	rows.Scan(arr...)
	//	users = append(users, user)
	//	suite.T().Log(user)
	//}

	rows.Close()

	// 方案3 xdb接口
	//err = suite.r.Select(&users, querySql)
	//suite.Nil(err)
	suite.T().Log(users)

}
