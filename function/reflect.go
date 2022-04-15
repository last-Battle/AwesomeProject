package function

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println("type:%v\n", v)
}

func Reflect() {
	var a float32 = 3.14
	reflectType(a)
}
