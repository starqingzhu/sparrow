package rf

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRf(t *testing.T) {
	//type cat struct {
	//}
	//ins := &cat{}
	//typeofCat := reflect.TypeOf(ins)
	//log.Println("name:", typeofCat.Name(), "kind:", typeofCat.Kind())
	//
	//typeofCat = typeofCat.Elem()
	//log.Println("Elem name:", typeofCat.Name(), "Elem kind:", typeofCat.Kind())

	//获取结构体类型成员类型

	type StructField struct {
		Name    string `json:"type" id:"100"` // 字段名
		PkgPath string // 字段路径
		//Type      Type      // 字段反射类型对象
		//Tag       StructTag // 字段的结构体标签
		Offset    uintptr // 字段在结构体中的相对偏移
		Index     []int   // Type.FieldByIndex中的返回的索引值
		Anonymous bool    // 是否为匿名字段
	}

	var stf = StructField{
		Name:    "111",
		PkgPath: "xxxx",
	}

	valueofStf := reflect.ValueOf(&stf)
	value := valueofStf.Elem()
	typefofStf := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldInfo := value.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typefofStf.Field(i).Name, fieldInfo.Kind(), fieldInfo.Interface())

	}

	value.Field(0).SetString("fb_1")
	value.Field(1).SetString("fb_1_Pwd")

	for i := 0; i < value.NumField(); i++ {
		fieldInfo := value.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typefofStf.Field(i).Name, fieldInfo.Kind(), fieldInfo.Interface())

	}

	//typeofStf := reflect.TypeOf(stf)
	//log.Println("name:", typeofStf.Name(), "kind:", typeofStf.Kind())
	//
	//for i := 0; i < typeofStf.NumField(); i++ {
	//	field := typeofStf.Field(i)
	//	fmt.Printf("name: %v  tag: '%v'  kind:%v\n", field.Name, field.Tag, field.Type.Kind())
	//}
}
