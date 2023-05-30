package xdb

//func (suite *XDBTestSuite) TestXdbUserQuery() {
//
//	querySql := "select * from user"
//	var users []tbstruct.User
//
//	rows, err := suite.r.DB.Query(querySql)
//	suite.Nil(err)
//
//	//rows.Scan(&users)
//	for rows.Next() {
//		var user tbstruct.User
//		arr := getObjMemAddrArr(&user)
//		rows.Scan(arr...)
//		users = append(users, user)
//		suite.T().Log(user)
//	}
//	rows.Close()
//
//	suite.r.DB.Select(&users, querySql)
//
//}
//
//func getObjMemAddrArr(p *tbstruct.User) []interface{} {
//
//	ret := make([]interface{}, 0)
//	value := reflect.ValueOf(p)
//	typ := reflect.TypeOf(p)
//
//	for i := 0; i < typ.NumField(); i++ {
//		//field := typ.Field(i)
//		filedValue := value.Field(i)
//
//		ret = append(ret, filedValue.UnsafePointer())
//	}
//
//	return nil
//}
