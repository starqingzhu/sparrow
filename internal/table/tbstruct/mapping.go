package tbstruct

import (
	"errors"
	"fmt"
	"reflect"
)

/*
映射关系
*/

var stName2DbTbName = map[string]string{
	"User": "user",
}

func GetStructTbName(st interface{}) (string, error) {
	stName := reflect.TypeOf(st).Name()
	return GetTbName(stName)
}

func GetTbName(stName string) (string, error) {
	var tbName, ok = stName2DbTbName[stName]
	if !ok {
		return "", errors.New(fmt.Sprintf("%s not match tbname", stName))
	}

	return tbName, nil
}
