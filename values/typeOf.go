package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.142))
	fmt.Println(reflect.TypeOf(string(65))) // A string
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hello"))
}
