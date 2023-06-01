package tbstruct

import (
	"errors"
	"fmt"
	"reflect"
)

/*
功能： 输入结构体对象指针，返回对象成员的 指针
缺陷：只支持一级结构体，不支持嵌套
*/
func GetObjMemAddrArr(structure interface{}) ([]interface{}, error) {

	t := reflect.ValueOf(structure)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New(fmt.Sprintf("p:%v, is not ptr", structure))
	}

	value := t.Elem()
	if value.Kind() != reflect.Struct {
		return nil, errors.New(fmt.Sprintf("expected a struct, got %s", value.Kind()))
	}

	ret := make([]interface{}, value.NumField())

	for i := 0; i < value.NumField(); i++ {
		//获取成员值和类型
		fieldValue := value.Field(i)

		//获取成员类型的地址指针
		typePtr := fieldValue.Addr().Interface()
		//typePtr := unsafe.Pointer(&fieldType.Type)
		fmt.Printf("Field %d: %+v %v, kind:%v\n", i, fieldValue.Interface(), typePtr, fieldValue.Kind())
		//addr := unsafe.Pointer(value.FieldByName(name).UnsafeAddr())
		ret = append(ret, typePtr)
	}

	return ret, nil
}
