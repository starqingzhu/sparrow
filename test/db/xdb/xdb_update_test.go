package xdb

import "sparrow/internal/table/tbstruct"

func (suite *XDBTestSuite) TestXdbUserUpdate() {
	data := tbstruct.User{
		Id:   1700635576010,
		Name: "update111",
		Age:  10,
	}

	tbName, err := tbstruct.GetStructTbName(data)
	suite.Nil(err)
	if err != nil {
		return
	}
	var condition = map[string]interface{}{
		"Id =": 1700635576010,
	}

	suite.r.Update(tbName, data, condition)
}
