package xdb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

/*
数据库基本操作
增删改查
*/
func generateInsertSql(data interface{}, tbName string) (string, []interface{}) {
	//获取结构体类型信息
	typ := reflect.TypeOf(data).Elem()

	//获取结构体字段数量
	numFields := typ.NumField()

	//构建sql
	insertSQL := fmt.Sprintf("INSERT INTO %s(", tbName)
	valuesSQL := " VALUES("
	var values []interface{}
	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		tagDbName := field.Tag.Get("db")
		insertSQL += tagDbName + ","
		valuesSQL += "?,"

		fieldName := field.Name
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface()
		values = append(values, fieldValue)
	}

	insertSQL = insertSQL[:len(insertSQL)-1] + ")"
	valuesSQL = valuesSQL[:len(valuesSQL)-1] + ")"
	insertSQL += valuesSQL

	return insertSQL, values

}

func generateDeleteSql(tbName string, condition map[string]interface{}) string {
	var conditions []string

	for field, value := range condition {
		conditions = append(conditions, fmt.Sprintf("%s %v", field, value))
	}

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s;", tbName, strings.Join(conditions, " AND "))

	return sql
}

func generateUpdateSql(tbName string, st interface{}, condition map[string]interface{}) string {
	typ := reflect.TypeOf(st)

	var updates []string
	var conditions []string

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldVale := reflect.ValueOf(st).Field(i).Interface()

		updates = append(updates, fmt.Sprintf("%s = '%v'", fieldName, fieldVale))
	}

	for field, value := range condition {
		conditions = append(conditions, fmt.Sprintf("%s %v", field, value))
	}

	updateSQL := fmt.Sprintf("UPDATE %s SET WHERE %s;", tbName, strings.Join(updates, ", "), strings.Join(conditions, " AND "))

	return updateSQL
}

func generateSelectSql(tbName string, condition map[string]interface{}, limit int64) string {
	var conditions []string
	for field, value := range condition {
		conditions = append(conditions, fmt.Sprintf("%s '%v'", field, value))
	}

	var selectSQL = ""
	if limit > 0 {
		selectSQL = fmt.Sprintf("SELECT * FROM %s WHERE %s limit %d;", tbName, strings.Join(conditions, " AND "), limit)
	} else {
		selectSQL = fmt.Sprintf("SELECT * FROM %s WHERE %s;", tbName, strings.Join(conditions, " AND "))
	}

	return selectSQL
}

func (db *XDB) Insert(tbName string, st interface{}) error {
	//根据结构生成sql
	sql, vals := generateInsertSql(st, tbName)
	//调用sql
	_, err := db.DB.Exec(sql, vals...)
	return err
}

/*
参数：
tbName 表名
condition 条件  key:表达式 eg(Id<)  value：值或者表达式
*/
func (db *XDB) Delete(tbName string, condition map[string]interface{}) error {
	//生成sql
	sql := generateDeleteSql(tbName, condition)
	//调用sql
	_, err := db.DB.Exec(sql)
	return err
}

func (db *XDB) Update(tbName string, st interface{}, condition map[string]interface{}) error {
	//生成sql
	sql := generateUpdateSql(tbName, st, condition)
	//调用sql
	_, err := db.DB.Exec(sql)
	return err
}

func (db *XDB) Select(tbName string, condition map[string]interface{}, limit int64) (*sqlx.Rows, error) {
	//var ret = make([]interface{}, 0)
	//生成sql
	sql := generateSelectSql(tbName, condition, limit)
	//调用sql
	rows, err := db.DB.Queryx(sql)
	if err != nil {
		return nil, err
	}

	return rows, err
}
