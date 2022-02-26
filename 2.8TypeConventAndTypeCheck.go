package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}
