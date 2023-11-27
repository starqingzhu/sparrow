package xdb

import (
	"fmt"
	"sparrow/internal/table/tbstruct"
	"time"
)

func (suite *XDBTestSuite) TestXdbUserInsert() {
	//var count = 10
	//for i := 0; i < count; i++ {
	//	name := fmt.Sprintf("name%d", i)
	//	result, err := suite.r.Exec("INSERT INTO user(id,name,age) VALUES(?,?,?)", i, name, i)
	//	suite.Nil(err)
	//	log.Println(result)
	//}
	var nowTm = time.Now().UnixMilli()

	data := tbstruct.User{
		Id:   nowTm,
		Name: fmt.Sprintf("name:%d", nowTm),
		Age:  int(nowTm) % 80,
	}
	tbName, err := tbstruct.GetStructTbName(data)
	suite.Nil(err)
	if err != nil {
		return
	}

	suite.r.Insert(tbName, &data)
}
