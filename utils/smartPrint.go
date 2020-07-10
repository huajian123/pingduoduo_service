package utils

import (
	"fmt"
	"reflect"
)

func SmartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获得的数据")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}
