package xdb

import (
	"fmt"
	"log"
)

func (suite *XDBTestSuite) TestXdbUserInsert() {
	var count = 10
	for i := 0; i < count; i++ {
		var i = 1
		name := fmt.Sprintf("name%d", i)
		result, err := suite.r.Exec("INSERT INTO user(id,name,age) VALUES(?,?,?)", i, name, i)
		suite.Nil(err)
		log.Println(result)
	}
}
