package xdb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sparrow/pkg/log/zaplog"
)

type (
	XDbConf struct {
		User         string `json:"user"`
		Password     string `json:"password"`
		ServerIp     string `json:"serverip"`
		Port         int32  `json:"port"`
		DataBase     string `json:"database"`
		MaxOpenConns int64  `json:"maxopenconns"`
		MaxIdleConns int64  `json:"maxidleconns"`
	}

	XDB struct {
		*sqlx.DB
	}
)

var Main *XDB

func init() {
	Main = nil
}

func Init(conf *XDbConf) (*XDB, error) {
	if conf == nil {
		errStr := "xdb init error, conf is nil"
		zaplog.LoggerSugar.Errorf(errStr)
		return nil, errors.New(errStr)
	}
	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, "link", conf.ServerIp, conf.Port, conf.DataBase)

	// 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		zaplog.LoggerSugar.Errorf("xdb init error, err:%s, conf:%v", err.Error(), *conf)
		return nil, err
	}

	// 设置连接池参数
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	Main = &XDB{db}

	return Main, nil
}

func (db *XDB) CloseWrapper() {
	if db != nil {
		db.Close()
	}
}

// 创建表
//func (db *XDB) CreateTable(tb interface{}) error {
//	err := xdb2.CreateTableWithDB(db, tb)
//	return err
//}

// QueryRow 封装查询单行数据的方法
func (db *XDB) QueryRow(query string, args ...interface{}) *sqlx.Row {
	return db.DB.QueryRowx(query, args...)
}

// Query 封装查询多行数据的方法
func (db *XDB) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.DB.Queryx(query, args...)
	if err != nil {
		zaplog.LoggerSugar.Errorf("xdb query error, err:%s, query:%s, args:%v", err.Error(), query, args)
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	columns, _ := rows.Columns()
	for rows.Next() {
		row := make(map[string]interface{})
		dest := make([]interface{}, len(columns))
		for i := range dest {
			dest[i] = &dest[i]
		}

		if err = rows.Scan(dest...); err != nil {
			zaplog.LoggerSugar.Errorf("xdb scan error, err:%s, query:%s, args:%v, dest:%v", err.Error(),
				query, args, dest)
			return nil, err
		}
		for i, column := range columns {
			row[column] = dest[i].(interface{}).(*string)
		}
		result = append(result, row)
	}

	return result, err
}

// Exec 封装执行 SQL 语句的方法
func (db *XDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}
