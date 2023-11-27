package xdb

import (
	"sparrow/internal/table/tbstruct"
)

func (suite *XDBTestSuite) TestXdbUserDelete() {
	data := tbstruct.User{}

	tbName, err := tbstruct.GetStructTbName(data)
	suite.Nil(err)
	if err != nil {
		return
	}
	var condition = map[string]interface{}{
		"Id <": 1000,
	}

	suite.r.Delete(tbName, condition)
}
