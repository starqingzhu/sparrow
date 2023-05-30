package xdb

import "sparrow/pkg/log/zaplog"

func (suite *XDBTestSuite) TestCreateTable() {
	zaplog.LoggerSugar.Info("test....")

	// 建表语句
	createTableSql := `CREATE TABLE IF NOT EXISTS user 
					   (id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
						name VARCHAR(50) NOT NULL,
						age INT(11) NOT NULL,
						PRIMARY KEY (id));`
	suite.r.DB.MustExec(createTableSql)
}
