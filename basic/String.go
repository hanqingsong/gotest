package main

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

func main() {
	value := 3
	fmt.Println(reflect.TypeOf(value))
	valueStr := cast.ToString(value)
	fmt.Println(reflect.TypeOf(valueStr))

}
